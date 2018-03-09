package badger

import (
	"log"

	// Avoid naming conflicts
	bdg "github.com/dgraph-io/badger"
)

// Store implements the StreamStore interface against a badger backend.
type Store struct {
	*bdg.DB
}

// New creates a new badger client
func New() (b *Store, err error) {
	opts := bdg.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"

	db, err := bdg.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	b = &Store{db}
	b.streamInit()

	return b, nil
}
