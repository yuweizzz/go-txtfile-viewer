package pkg

import (
	"bytes"
	"embed"
	"text/template"
	"net/http"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

//go:embed static/html.tpl
var tplfile embed.FS

type PageData struct {
	Content string
	Title   string
}

func (p *PageData) Pretty() {
	c := strings.Split(p.Content, "\n")
	for num, value := range(c) {
		if value != ""{
			c[num] = "<br/><p>" + value + "</p>"
		}
	}
	p.Content = strings.Join(c, "")
}

func RenderPage(file http.File, filename string) (rs *bytes.Reader) {
	buf := &bytes.Buffer{}
	buf.ReadFrom(file)
	pd := &PageData{
		Content: buf.String(),
		Title:   filename,
	}
	if strings.HasSuffix(filename, ".md") {
		md := goldmark.New(goldmark.WithExtensions(extension.GFM))
		buf.Reset()
		md.Convert([]byte(pd.Content), buf)
		pd.Content = buf.String()
		buf.Reset()
		tpl := template.Must(template.New("html.tpl").ParseFS(tplfile, "static/html.tpl"))
		if err := tpl.ExecuteTemplate(buf, "html.tpl", pd); err != nil {
			panic(err)
		}
		rs = bytes.NewReader(buf.Bytes())
		return
	}
	pd.Pretty()
	buf.Reset()
	tpl := template.Must(template.New("html.tpl").ParseFS(tplfile, "static/html.tpl"))
	if err := tpl.ExecuteTemplate(buf, "html.tpl", pd); err != nil {
		panic(err)
	}
	rs = bytes.NewReader(buf.Bytes())
	return
}

func RenderBuffer(buf *bytes.Buffer) *bytes.Buffer {
	pd := &PageData{
		Content: buf.String(),
		Title:   "a",
	}
	pd.Pretty()
	buf.Reset()
	tpl := template.Must(template.New("html.tpl").ParseFS(tplfile, "static/html.tpl"))
	if err := tpl.ExecuteTemplate(buf, "html.tpl", pd); err != nil {
		panic(err)
	}
	return buf
}
