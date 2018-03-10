package main

import (
	"log"
	"path/filepath"

	"github.com/chulabs/seer/server"
)

var path = filepath.FromSlash("/var/seer")

func main() {
	srv, err := server.New(path)
	if err != nil {
		log.Fatal(path)
		return
	}

}
