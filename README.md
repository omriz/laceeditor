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