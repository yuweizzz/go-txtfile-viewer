package pkg

import (
	"io/fs"
	"net/http"
	"strings"
)

type TextFile struct {
	http.File
}

func (f TextFile) Readdir(n int) (fis []fs.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files {
		filename := file.Name()
		if !strings.HasPrefix(filename, ".") {
			if file.IsDir() {
				fis = append(fis, file)
			} else if strings.HasSuffix(filename, ".txt") {
				fis = append(fis, file)
			} else if strings.HasSuffix(filename, ".md") {
				fis = append(fis, file)
			} else if strings.HasSuffix(filename, ".markdown") {
				fis = append(fis, file)
			}
		}
	}
	return
}

type TextFileSystem struct {
	http.FileSystem
}

func (fsys TextFileSystem) Open(name string) (http.File, error) {
	file, err := fsys.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return TextFile{file}, err
}
