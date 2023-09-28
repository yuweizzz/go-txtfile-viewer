package pkg

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"
)

//go:embed static/html.tpl
var tplfile embed.FS

type PageData struct {
	Content string
	Title   string
}

func RenderPage(file http.File, filename string) (rs *bytes.Reader) {
	buf := &bytes.Buffer{}
	buf.ReadFrom(file)
	pd := PageData{
		Content: buf.String(),
		Title:   filename,
	}
	buf.Reset()
	tpl := template.Must(template.New("html.tpl").ParseFS(tplfile, "static/html.tpl"))
	if err := tpl.ExecuteTemplate(buf, "html.tpl", pd); err != nil {
		panic(err)
	}
	rs = bytes.NewReader(buf.Bytes())
	return
}
