package pkg

import (
	"bytes"
	"crypto/sha1"
	"embed"
	"encoding/hex"
	"io"
	"net/http"
	"strings"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

//go:embed static/html.tpl
var tplfile embed.FS

type FileType int

const (
	Raw FileType = iota
	Txt
	Markdown
)

type PageData struct {
	File     http.File
	Content  string // mark: need to refactor
	Title    string
	Type     FileType
	CheckSum string
}

func NewPageData(file http.File) *PageData {
	stat, _ := file.Stat()
	pd := &PageData{
		Title: stat.Name(),
		File:  file,
	}
	if strings.HasSuffix(pd.Title, ".txt") {
		pd.Type = Txt
	} else {
		pd.Type = Markdown
	}
	return pd
}

func (pd *PageData) SumContent() {
	raw, _ := io.ReadAll(pd.File)
	h := sha1.New()
	// mark: need to refactor
	pd.Content = string(raw)
	io.WriteString(h, pd.Content)
	pd.CheckSum = hex.EncodeToString(h.Sum(nil)[:])
}

func (pd *PageData) Pretty() {
	switch pd.Type {
	case Txt:
		c := strings.Split(pd.Content, "\n")
		for num, value := range c {
			if value != "" {
				c[num] = "<p>" + value + "</p>"
			}
		}
		pd.Content = strings.Join(c, "")
	case Markdown:
		buf := &bytes.Buffer{}
		md := goldmark.New(goldmark.WithExtensions(extension.GFM))
		md.Convert([]byte(pd.Content), buf)
		pd.Content = buf.String()
	default:
		pd.Content = "<pre>\n" + pd.Content + "</pre>\n"
	}
}

func (pd *PageData) RenderPage() *bytes.Reader {
	buf := &bytes.Buffer{}
	tpl := template.Must(template.New("html.tpl").ParseFS(tplfile, "static/html.tpl"))
	if err := tpl.ExecuteTemplate(buf, "html.tpl", pd); err != nil {
		panic(err)
	}
	return bytes.NewReader(buf.Bytes())
}

func RenderBuffer(title string, buf *bytes.Buffer) *bytes.Buffer {
	pd := &PageData{
		Content: buf.String(),
		Title:   title,
	}
	pd.Pretty()
	buf.Reset()
	tpl := template.Must(template.New("html.tpl").ParseFS(tplfile, "static/html.tpl"))
	if err := tpl.ExecuteTemplate(buf, "html.tpl", pd); err != nil {
		panic(err)
	}
	return buf
}
