package bolt

import (
	"github.com/chulabs/seer/store"
	"github.com/chulabs/seer/stream"
	"github.com/vmihailenco/msgpack"

	// Avoid namespace conflicts
	blt "github.com/boltdb/bolt"
)

// streamBucket is the key for the stream bucket.
var streamBucket = []byte("streams")

// CreateStream saves the provided stream at name, returns an error if a
// stream already exists at that address.
func (b *Store) CreateStream(name string, s *stream.Stream) (err error) {
	err = b.Update(func(tx *blt.Tx) error {
		bk := tx.Bucket(streamBucket)

		val := bk.Get([]byte(name))
		if val != nil {
			return &store.AlreadyExistsError{Kind: "stream", Entity: name}
		}

		val, _ = msgpack.Marshal(s)
		err := bk.Put([]byte(name), val)
		return err
	})

	return err
}

// GetStream returns the stream stored at name, or an error if the stream does
// not exist, or has corrupted data.
func (b *Store) GetStream(name string) (s *stream.Stream, err error) {
	s = &stream.Stream{}

	err = b.View(func(tx *blt.Tx) error {
		bk := tx.Bucket(streamBucket)

		val := bk.Get([]byte(name))
		if val == nil {
			return &store.NotFoundError{Kind: "stream", Entity: name}
		}
		err = msgpack.Unmarshal(val, s)
		return err
	})

	if err != nil {
		return nil, err
	}
	return s, nil
}

// DeleteStream deletes the stream stored at name, or returns an error if
// no such stream exists.
func (b *Store) DeleteStream(name string) (err error) {
	err = b.Update(func(tx *blt.Tx) error {
		bk := tx.Bucket(streamBucket)

		val := bk.Get([]byte(name))
		if val == nil {
			return &store.NotFoundError{Kind: "stream", Entity: name}
		}

		err := bk.Delete([]byte(name))
		return err
	})

	return err
}

// ListStreams returns a paged list of streams, or an error if none are found.
func (b *Store) ListStreams(pageNum, pageSize int) (s []*stream.Stream, err error) {
	return
}

// UpdateStream overwrites the stream at name with the provided stream, or
// returns an error if no stream exists at name.
func (b *Store) UpdateStream(name string, s *stream.Stream) (err error) {
	return
}