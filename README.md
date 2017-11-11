# LAce Editor (i.e, Local ACE)

This is a personal project trying to solve a simple use case:
You are working on a remote VM, say from your chromebook, and you have SSH
access. You'd like to edit some code you are working on, so you either access it
with your own local computer and sync to the remote machine or edit it on the
remote machine using VIM or some other console editor.
LAce comes to be something in between, it will launch ACE editor with a local
WebDav backend that you can access from a browser giving you the look and feel
of an IDE, and can be more resiliant to network disruptions.

## General overview:
1. Backend using webdav - either https://godoc.org/golang.org/x/net/webdav or https://www.npmjs.com/package/webdav-server
2. Frontend using ACE editor (https://github.com/ajaxorg/ace-builds/) and probably a webdav library (https://www.npmjs.com/package/webdav)
3. We'll wrap the frontend with some simple UI framework like https://getmdl.io or maybe something more soffisticated like Reactjs
4. Using https://github.com/GeertJohan/go.rice to serve static files. We can use a go:generate command to compile our static files beforehand.

## How to build/install
Assming the standard go project structure:
```bash
# Fetching go.rice
go get github.com/GeertJohan/go.rice
go get github.com/GeertJohan/go.rice/rice
# Fetching laceeditor
go get github.com/omriz/laceeditor
# Building static sources
cd src/github.com/omriz/laceeditor
../../../../bin/rice embed-go
cd -
# Installing LAce
go install github.com/omriz/laceeditor
```