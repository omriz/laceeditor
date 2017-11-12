package main

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"golang.org/x/net/webdav"
	"net/http"
)

// Assuming you did go get github.com/GeertJohan/go.rice/rice
//go:generate $GOPATH/bin/rice embed-go

var davDir = flag.String("dav_dir", "/tmp", "Directory where the files will be accessible")

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
	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	// In order to bind to localhost we need to do
	// http.ListenAndServe("localhost:8080", nil)
	http.ListenAndServe(":8080", nil)
}
