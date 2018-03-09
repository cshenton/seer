package bolt

import (
	// Avoid namespace conflicts
	blt "github.com/boltdb/bolt"
)

// Store wraps a bolt DB and fulfills the store.StreamStore interface.
type Store struct {
	*blt.DB
}

// New creates a bolt Store at the provided file path.
func New(path string) (b *Store, err error) {
	db, err := blt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	b = &Store{db}
	return b, nil
}
