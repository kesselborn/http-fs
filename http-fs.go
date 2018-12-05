// httpserver.go
// hello
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

/////////////////////////// 404 hack

type custom404 struct {
	http.ResponseWriter
	contentType string
	content     string
	is404       *bool
}

func (w custom404) WriteHeader(code int) {
	if code == 404 {
		w.Header().Set("Content-Type", w.contentType)
		*w.is404 = true
	}
	w.ResponseWriter.WriteHeader(code)
}

func (w custom404) Write(b []byte) (int, error) {
	if *w.is404 {
		w.ResponseWriter.Write([]byte(w.content))
		return len(b), nil
	}

	return w.ResponseWriter.Write(b)
}

/////////////////////////// 404 hack end

func serve(dir string, readOnly bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errorOut := func(format string, a ...interface{}) {
			msg := fmt.Sprintf(format, a...)
			log.Printf(msg)
			http.Error(w, msg, http.StatusInternalServerError)
		}

		logMsg := func(method string) {
			log.Printf("%-60s| header: %s\n", method+" "+r.URL.Path, r.Header)
		}

		if readOnly && r.Method != "GET" {
			http.Error(w, "http-fs was started in read-only mode", 405)
			return
		}

		switch r.Method {
		case "GET":
			logMsg("GET")
			http.FileServer(http.Dir(dir)).ServeHTTP(custom404{ResponseWriter: w, contentType: "text/html; charset: utf-8", content: custom404Content, is404: new(bool)}, r)
		case "DELETE":
			logMsg("DELETE")
			fileOrDir := path.Join(dir, r.URL.Path)
			if err := os.RemoveAll(fileOrDir); err != nil {
				errorOut("error removing %s: %s", fileOrDir, err)
				return
			}

			http.Error(w, http.StatusText(204), 204)
		case "PUT":
			logMsg("PUT")
			dir := dir + path.Dir(r.URL.Path)
			if err := os.MkdirAll(dir, 0755); err != nil {
				errorOut("error creating dir %s: %s", dir, err)
				return
			}

			f, err := ioutil.TempFile(dir, "upload-blob.")
			if err != nil {
				errorOut("error creating dir temporary file: %s", err)
				return
			}
			defer f.Close()

			if _, err := io.Copy(f, r.Body); err != nil {
				errorOut("error storing file %s: %s", f.Name, err)
				return
			}

			if err := os.Rename(f.Name(), dir+r.URL.Path); err != nil {
				errorOut("error storing file %s: %s", f.Name(), err)
				return
			}

			fmt.Fprintf(w, http.StatusText(201)+"\n")
		}
	}
}

func main() {
	addr_param := flag.String("addr", "0.0.0.0:5000", "where to listen for connection")
	dir := flag.String("dir", ".", "which directory to take as a root")
	read_only := flag.Bool("read-only", false, "start server read only")
	// hallo

	flag.Parse()
	fmt.Printf(`
serving current %s at %s

- get file %s/foo/bar:

    curl -O %s/foo/bar

- delete file %s/foo/bar (not available in read-only mode):

    curl -XDELETE %s/foo/bar

- upload file '/tmp/file' to %s/foo/bar: (not available in read-only mode):

    curl -T /tmp/file %s/foo/bar

`, *dir, *addr_param,
		*dir, *addr_param,
		*dir, *addr_param,
		*dir, *addr_param)

	http.HandleFunc("/", serve(*dir, *read_only))
	panic(http.ListenAndServe(*addr_param, nil))
}
