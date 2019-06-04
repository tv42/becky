package main

import (
	"log"
	"net"
	"net/http"
)

//go:generate go run github.com/tv42/becky index.html

type HTML struct {
	asset
}

func html(a asset) HTML {
	return HTML{a}
}

func main() {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serving at http://%s/", l.Addr())
	http.Handle("/", index)
	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}
}
