// +build !dev

package main

import (
	"encoding/base64"
	"hash/fnv"
	"io"
	"net/http"
	"strings"
	"time"
)

type asset struct {
	Name    string
	Content string
	etag    string
}

func (a asset) init() asset {
	// This is a method to minize the namespace pollution. Use
	// chaining to make it callable in variable declarations.
	h := fnv.New64a()
	_, _ = io.WriteString(h, a.Content)
	a.etag = `"` + base64.StdEncoding.EncodeToString(h.Sum(nil)) + `"`
	return a
}

func (a asset) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if a.etag != "" && w.Header().Get("ETag") == "" {
		w.Header().Set("ETag", a.etag)
	}
	body := strings.NewReader(a.Content)
	http.ServeContent(w, req, a.Name, time.Time{}, body)
}
