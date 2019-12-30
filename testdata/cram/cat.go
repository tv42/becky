package main

import (
	"io"
	"log"
	"os"
)

//go:generate go run github.com/tv42/becky greeting.txt

func txt(a asset) string {
	return a.Content
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("cat.go: ")

	if _, err := io.WriteString(os.Stdout, greeting); err != nil {
		log.Fatal(err)
	}
}
