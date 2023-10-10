package pkg

import (
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"io"
	"bytes"
	"log"
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

func TextFileServer(root TextFileSystem) http.Handler {
	return &TextFileHandler{root}
}

type TextFileHandler struct {
	root TextFileSystem
}

func (f *TextFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	log.Println("Incoming Request:", upath)
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}

	file, err := f.root.Open(upath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if stat, _ := file.Stat(); !stat.IsDir() {
		rs := RenderPage(file, stat.Name())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeContent(w, r, upath, stat.ModTime(), rs)
		return
	}
//
//	if checkIfModifiedSince(r, stat.ModTime()) == condFalse {
//		writeNotModified(w)
//		return
//	}
//	setLastModified(w, d.ModTime())
	dirList(w, r, file.(http.File))
}

type anyDirs interface {
	len() int
	name(i int) string
	isDir(i int) bool
}

type fileInfoDirs []fs.FileInfo

func (d fileInfoDirs) len() int          { return len(d) }
func (d fileInfoDirs) isDir(i int) bool  { return d[i].IsDir() }
func (d fileInfoDirs) name(i int) string { return d[i].Name() }

type dirEntryDirs []fs.DirEntry

func (d dirEntryDirs) len() int          { return len(d) }
func (d dirEntryDirs) isDir(i int) bool  { return d[i].IsDir() }
func (d dirEntryDirs) name(i int) string { return d[i].Name() }

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

func dirList(w http.ResponseWriter, r *http.Request, f http.File) {
	// Prefer to use ReadDir instead of Readdir,
	// because the former doesn't require calling
	// Stat on every entry of a directory on Unix.
	var dirs anyDirs
	var err error
	if d, ok := f.(fs.ReadDirFile); ok {
		var list dirEntryDirs
		list, err = d.ReadDir(-1)
		dirs = list
	} else {
		var list fileInfoDirs
		list, err = f.Readdir(-1)
		dirs = list
	}

	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs.name(i) < dirs.name(j) })

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := bytes.NewBufferString("")
	for i, n := 0, dirs.len(); i < n; i++ {
		name := dirs.name(i)
		if dirs.isDir(i) {
			name += "/"
		}
		// name may contain '?' or '#', which must be escaped to remain
		// part of the URL path, and not indicate the start of a query
		// string or fragment.
		url := url.URL{Path: name}
		fmt.Fprintf(buf, "<a href=\"%s\">%s</a>\n", url.String(), htmlReplacer.Replace(name))
	}
	buf = RenderBuffer(buf)
	io.Copy(w, buf)
}

//
//func checkIfModifiedSince(r *Request, modtime time.Time) condResult {
//	if r.Method != "GET" && r.Method != "HEAD" {
//		return condNone
//	}
//	ims := r.Header.Get("If-Modified-Since")
//	if ims == "" || isZeroTime(modtime) {
//		return condNone
//	}
//	t, err := ParseTime(ims)
//	if err != nil {
//		return condNone
//	}
//	// The Last-Modified header truncates sub-second precision so
//	// the modtime needs to be truncated too.
//	modtime = modtime.Truncate(time.Second)
//	if ret := modtime.Compare(t); ret <= 0 {
//		return condFalse
//	}
//	return condTrue
//}
//
//
//
//func setLastModified(w ResponseWriter, modtime time.Time) {
//	if !isZeroTime(modtime) {
//		w.Header().Set("Last-Modified", modtime.UTC().Format(TimeFormat))
//	}
//}
//
//func writeNotModified(w ResponseWriter) {
//	// RFC 7232 section 4.1:
//	// a sender SHOULD NOT generate representation metadata other than the
//	// above listed fields unless said metadata exists for the purpose of
//	// guiding cache updates (e.g., Last-Modified might be useful if the
//	// response does not have an ETag field).
//	h := w.Header()
//	delete(h, "Content-Type")
//	delete(h, "Content-Length")
//	delete(h, "Content-Encoding")
//	if h.Get("Etag") != "" {
//		delete(h, "Last-Modified")
//	}
//	w.WriteHeader(StatusNotModified)
//}