// httpserver.go
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

type dirServer string

func (s dirServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errorOut := func(format string, a ...interface{}) {
		msg := fmt.Sprintf(format, a...)
		log.Printf(msg)
		http.Error(w, msg, http.StatusInternalServerError)
	}

	fmt.Printf("%s: got request to %s\n", time.Now(), r.URL.Path)
	switch r.Method {
	case "GET":
		http.FileServer(http.Dir(s)).ServeHTTP(w, r)
	case "DELETE":
		if err := os.Remove(string(s) + r.URL.Path); err != nil {
			errorOut("error removing %s: %s", string(s)+r.URL.Path, err)
			return
		}

		fmt.Fprintf(w, http.StatusText(204))
	case "PUT":
		if err := os.MkdirAll(string(s)+path.Dir(r.URL.Path), 0755); err != nil {
			errorOut("error creating dir %s: %s", path.Dir(r.URL.Path), err)
			return
		}
		f, err := os.Create(string(s) + r.URL.Path)
		if err != nil {
			errorOut("error creating dir %s: %s", path.Dir(r.URL.Path), err)
			return
		}
		defer f.Close()

		if _, err := io.Copy(f, r.Body); err != nil {
			errorOut("error storing file %s: %s", f.Name, err)
			return
		}

		fmt.Fprintf(w, http.StatusText(201))
	}
}

func main() {
	portParam := flag.String("addr", "localhost:8080", "where to listen for connection")
	dirParam := flag.String("dir", ".", "which directory to take as a root")
	flag.Parse()
	fmt.Printf(`
serving current %s at %s

- get file %s/foo/bar:

    curl -O %s%s/foo/bar

- delete file %s/foo/bar:

    curl -XDELETE %s%s/foo/bar

- upload file to %s/foo/bar:

    curl -T /tmp/file %s%s/foo/bar

`, *dirParam, *portParam,
		*dirParam, *portParam, *dirParam,
		*dirParam, *portParam, *dirParam,
		*dirParam, *portParam, *dirParam)

	dir := dirServer(*dirParam)

	panic(http.ListenAndServe(*portParam, dir))
}
