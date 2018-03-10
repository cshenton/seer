package server

import (
	"github.com/chulabs/seer/store"
	"github.com/chulabs/seer/store/bolt"
)

// Server fulfills the protocol buffer's SeerServer interface.
type Server struct {
	DB store.StreamStore
}

// New creates a database connection and returns a Server.
func New(path string) (srv *Server, err error) {
	db, err := bolt.New(path)
	if err != nil {
		return nil, err
	}
	return &Server{DB: db}, nil
}
