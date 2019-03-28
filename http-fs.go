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
	"strings"
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

func serve(dir string, readOnly bool, basicAuth string) func(w http.ResponseWriter, r *http.Request) {
	protected := false
	baUsername := ""
	baPassword := ""

	if basicAuthTokens := strings.Split(basicAuth, ":"); len(basicAuthTokens) == 2 {
		protected = true
		baUsername = basicAuthTokens[0]
		baPassword = basicAuthTokens[1]
	}

	log.Printf("protected mode: %t\n", protected)

	return func(w http.ResponseWriter, r *http.Request) {
		errorOut := func(format string, a ...interface{}) {
			msg := fmt.Sprintf(format, a...)
			log.Printf(msg)
			http.Error(w, msg, http.StatusInternalServerError)
		}

		username, password, basicAuthProvided := r.BasicAuth()
		r.SetBasicAuth(username, "***********")
		logMsg := func(method string) {
			log.Printf("%-60s| header: %s\n", method+" "+r.URL.Path, r.Header)
		}

		if readOnly && r.Method != "GET" {
			http.Error(w, "http-fs was started in read-only mode", 405)
			return
		}

		log.Printf("0: %s=%s / %s=%s / %t", username, baUsername, password, baPassword, basicAuthProvided)
		if protected {
			if !basicAuthProvided {
				http.Error(w, "authentication required", 401)
				return
			}

			log.Printf("%s=%s / %s=%s", username, baUsername, password, baPassword)
			if username != baUsername || password != baPassword {
				http.Error(w, "wrong username / password", 403)
				return
			}
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
			dir := path.Join(dir, path.Dir(r.URL.Path))
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

			if err := os.Rename(f.Name(), path.Join(dir, path.Base(r.URL.Path))); err != nil {
				errorOut("error storing file %s: %s", f.Name(), err)
				return
			}

			fmt.Fprintf(w, http.StatusText(201)+"\n")
		}
	}
}

func main() {
	addr := flag.String("addr", "0.0.0.0:5000", "where to listen for connection")
	dir := flag.String("dir", ".", "which directory to take as a root")
	readOnly := flag.Bool("read-only", false, "start server read only")
	basicAuth := flag.String("basic-auth", "", "set to something like 'username:password' for basic auth")
	tlsCert := flag.String("tls-cert", "", "https tls certificate (only necessary when running in https mode)")
	tlsKey := flag.String("tls-key", "", "https private key (only necessary when running in https mode)")

	flag.Parse()
	fmt.Printf(`
serving current %s at %s

- get file %s/foo/bar:

    curl -O %s/foo/bar

- delete file %s/foo/bar (not available in read-only mode):

    curl -XDELETE %s/foo/bar

- upload file '/tmp/file' to %s/foo/bar: (not available in read-only mode):

    curl -T /tmp/file %s/foo/bar

`, *dir, *addr,
		*dir, *addr,
		*dir, *addr,
		*dir, *addr)

	http.HandleFunc("/", serve(*dir, *readOnly, *basicAuth))
	if *tlsCert != "" && *tlsKey != "" {
		log.Fatal(http.ListenAndServeTLS(*addr, *tlsCert, *tlsKey, nil))
	} else {
		if *basicAuth != "" {
			log.Fatal("basic auth is only supported in tls mode")
		}
		log.Fatal(http.ListenAndServe(*addr, nil))
	}
}
