package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

//go:generate go run github.com/tv42/becky greeting.txt

func txt(a asset) asset {
	return a
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("fakehttp.go: ")

	rec := httptest.NewRecorder()
	rec.Body = &bytes.Buffer{}
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}
	greeting.ServeHTTP(rec, req)
	if rec.Code != 200 {
		log.Fatalf("bad response: %v: %s", rec.Code, rec.Body.Bytes())
	}
	if _, err := os.Stdout.Write(rec.Body.Bytes()); err != nil {
		log.Fatal(err)
	}
}
