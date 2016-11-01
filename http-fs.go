// httpserver.go
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

type dirServer string

func (s dirServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errorOut := func(format string, a ...interface{}) {
		msg := fmt.Sprintf(format, a...)
		log.Printf(msg)
		http.Error(w, msg, http.StatusInternalServerError)
	}

	logMsg := func(method string) {
		log.Printf("%-60s| header: %s\n", method+" "+r.URL.Path, r.Header)
	}

	switch r.Method {
	case "GET":
		logMsg("GET")
		http.FileServer(http.Dir(s)).ServeHTTP(w, r)
	case "DELETE":
		logMsg("DELETE")
		fileOrDir := string(s) + path.Dir(r.URL.Path)
		if err := os.RemoveAll(fileOrDir); err != nil {
			errorOut("error removing %s: %s", fileOrDir, err)
			return
		}

		fmt.Fprintf(w, http.StatusText(204)+"\n")
	case "PUT":
		logMsg("PUT")
		dir := string(s) + path.Dir(r.URL.Path)
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

		if err := os.Rename(f.Name(), string(s)+r.URL.Path); err != nil {
			errorOut("error storing file %s: %s", f.Name(), err)
			return
		}

		fmt.Fprintf(w, http.StatusText(201)+"\n")
	}
}

func main() {
	addrParam := flag.String("addr", "localhost:8080", "where to listen for connection")
	dirParam := flag.String("dir", ".", "which directory to take as a root")
	flag.Parse()
	fmt.Printf(`
serving current %s at %s

- get file %s/foo/bar:

    curl -O %s/foo/bar

- delete file %s/foo/bar:

    curl -XDELETE %s/foo/bar

- upload file '/tmp/file' to %s/foo/bar:

    curl -T /tmp/file %s/foo/bar

`, *dirParam, *addrParam,
		*dirParam, *addrParam,
		*dirParam, *addrParam,
		*dirParam, *addrParam)

	dir := dirServer(*dirParam)

	panic(http.ListenAndServe(*addrParam, dir))
}
