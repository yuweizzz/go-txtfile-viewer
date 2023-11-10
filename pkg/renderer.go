package pkg

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"embed"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/saintfish/chardet"
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
	Content  string
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

func (pd *PageData) SumContent() string {
	raw, _ := io.ReadAll(pd.File)
	sum := sha1.Sum(raw)
	pd.CheckSum = hex.EncodeToString(sum[:])
	detector := chardet.NewTextDetector()
	charset, err := detector.DetectBest(raw)
	if err != nil {
		panic(err)
	}
	pd.Content = string(raw)
	return charset.Charset
}

func (pd *PageData) Pretty() {
	buf := &bytes.Buffer{}
	switch pd.Type {
	case Txt:
		scanner := bufio.NewScanner(strings.NewReader(pd.Content))
		// MaxScanTokenSize: 64k 
		readbuf := make([]byte, 0, 64*1024)
		// Resize Buffer: 1M
		scanner.Buffer(readbuf, 1024*1024)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) != 0 {
				fmt.Fprintf(buf, "<p>%s</p>", line)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(buf, "scanning file error: ", err)
		}
		pd.Content = buf.String()
	case Markdown:
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
		return nil
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
		return nil
	}
	return buf
}
