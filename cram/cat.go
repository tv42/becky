package main

import (
	"io"
	"log"
	"os"
)

//go:generate -command asset go run asset.go
//go:generate asset greeting.txt

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
