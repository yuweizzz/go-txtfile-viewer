package pkg

import (
	"html/template"
	"embed"
	"net/http"
	"io"
	//"fmt"
	"bytes"
	//"bufio"
)

//go:embed static/html.tpl
var tplfile embed.FS

type PageData struct {
	Content []byte
	Title string
}

func RenderPage(file http.File, filename string) (rs *bytes.Reader) {
	data,_ := io.ReadAll(file)
	pd := &PageData{
		Content: data,
		Title: filename,
	}
	var buf bytes.Buffer
	tpl := template.Must(template.New("html.tpl").ParseFS(tplfile, "static/html.tpl"))
    if err := tpl.ExecuteTemplate(&buf, "html.tpl", pd); err != nil {
        panic(err)
    }
	rdbuf := make([]byte, 100000)
	_, err := buf.Read(rdbuf)
	if err != nil {
		panic(err)
	}
	c := bytes.NewReader(rdbuf)
	return c
}
