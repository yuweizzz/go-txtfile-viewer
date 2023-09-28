package pkg

import (
	"bytes"
	"embed"
	"text/template"
	"net/http"
	"strings"
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
	pd.Pretty()
	buf.Reset()
	tpl := template.Must(template.New("html.tpl").ParseFS(tplfile, "static/html.tpl"))
	if err := tpl.ExecuteTemplate(buf, "html.tpl", pd); err != nil {
		panic(err)
	}
	rs = bytes.NewReader(buf.Bytes())
	return
}
