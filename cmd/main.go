package main

import (
	"log"
	"net/http"

	"github.com/yuweizzz/go-txtfile-viewer/pkg"
)

func main() {
	fsys := pkg.CustomFileSystem{http.Dir("/tmp/text_files")}
	http.Handle("/", pkg.CustomFileServer(fsys))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
