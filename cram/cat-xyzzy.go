package main

import (
	"io"
	"log"
	"os"
)

func txt(a asset) string {
	return a.Content
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("cat-xyzzy.go: ")

	if _, err := io.WriteString(os.Stdout, xyzzy); err != nil {
		log.Fatal(err)
	}
}
