package main

import (
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	handler := &webdav.Handler{
		Prefix:     "/",
		FileSystem: webdav.Dir("./data"),
		LockSystem: webdav.NewMemLS(),
	}

	http.Handle("/", handler)

	http.ListenAndServe("0.0.0.0:8090", nil)
}
