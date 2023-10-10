package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/yuweizzz/go-txtfile-viewer/pkg"
)

var (
	Port int
	Dir  string
)

func init() {
	flag.IntVar(&Port, "p", 8080, "The listen port")
	flag.StringVar(&Dir, "d", ".", "The dir where to serve")
}

func main() {
	flag.Parse()
	fsys := pkg.TextFileSystem{http.Dir(Dir)}
	http.Handle("/", pkg.TextFileServer(fsys))
	http.ListenAndServe(":"+strconv.Itoa(Port), nil)
}
