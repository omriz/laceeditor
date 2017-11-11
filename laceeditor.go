package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/webdav"
	"net/http"
)

var davDir = flag.String("dav_dir", "/tmp", "Directory where the files will be accessible")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	flag.Parse()
	davUrl := "/_/dav/"
	dav := webdav.Handler{
		Prefix:     davUrl,
		FileSystem: webdav.Dir(*davDir),
		LockSystem: webdav.NewMemLS(),
		Logger:     nil,
	}
	http.HandleFunc(davUrl, dav.ServeHTTP)
	http.HandleFunc("/func/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
