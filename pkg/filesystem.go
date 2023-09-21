package pkg

import (
	"io/fs"
	"net/http"
	"strings"
)

func confirmTxtfile(name string) bool {
	return strings.HasSuffix(name, ".txt")
}

type txtFile struct {
	http.File
}

func (f txtFile) Readdir(n int) (fis []fs.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files {
		if confirmTxtfile(file.Name()) {
			fis = append(fis, file)
		}
	}
	return
}

type CustomFileSystem struct {
	http.FileSystem
}

func (fsys CustomFileSystem) Open(name string) (http.File, error) {
	file, err := fsys.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return txtFile{file}, err
}
