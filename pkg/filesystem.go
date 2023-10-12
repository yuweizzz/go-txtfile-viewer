package pkg

import (
	"io/fs"
	"net/http"
	"regexp"
	"strings"
)

type TextFile struct {
	http.File
}

var re = regexp.MustCompile(`^.*\.(txt|md|markdown)$`)

func (f TextFile) Readdir(n int) (fis []fs.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files {
		filename := file.Name()
		if !strings.HasPrefix(filename, ".") {
			if file.IsDir() || re.MatchString(filename) {
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
