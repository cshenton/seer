package main

import (
	"log"
	"net"
	"path/filepath"

	"github.com/chulabs/seer/seer"
	"github.com/chulabs/seer/server"
	"google.golang.org/grpc"
)

var port = ":8080"
var path = filepath.FromSlash("/var/seer")

func main() {
	srv, err := server.New(path)
	if err != nil {
		log.Fatal("failed to create server:", err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	s := grpc.NewServer()
	seer.RegisterSeerServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
