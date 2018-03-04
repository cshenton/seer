package badger

import (
	"log"

	"github.com/chulabs/seer/store"
	"github.com/chulabs/seer/stream"
	"github.com/vmihailenco/msgpack"
)

// Key for the stream index
var streamList = []byte("stream_list")

// Key for a particular stream's data
func streamKey(name string) []byte {
	return []byte("stream::" + name)
}

// CreateStream saves the provided stream at name, returns an error if a
// stream already exists at that address.
func (b *Store) CreateStream(name string, s *stream.Stream) (err error) {
	// Get streamlist
	// if name in streamList
	// 	abort tx, return error
	// add name to streamList
	// serialize s to []byte
	// Put streamList
	// Put s to stream::name
	// Commit tx
	return
}

// GetStream returns the stream stored at name, or an error if no such
// stream exists.
func (b *Store) GetStream(name string) (s *stream.Stream, err error) {
	tx := b.NewTransaction(false)

	item, err := tx.Get(streamKey(name))
	if err != nil {
		return nil, &store.NotFoundError{Kind: "stream", Entity: name}
	}

	s = &stream.Stream{}
	err = msgpack.Unmarshal(item.Key(), s)
	if err != nil {
		log.Print(string(item.Key()))
		err = &store.CorruptDataError{Kind: "stream"}
		return nil, err
	}

	return s, nil
}

// DeleteStream deletes the stream stored at name, or returns an error if
// no such stream exists.
func (b *Store) DeleteStream(name string) (err error) {
	// txn.Delete on stream::name, bubble up error.
	return
}

// ListStreams returns a paged list of streams, or an error if none are found.
func (b *Store) ListStreams(pageNum, pageSize int) (s []*stream.Stream, err error) {
	return
}

// UpdateStream overwrites the stream at name with the provided stream, or
// returns an error if no stream exists at name.
func (b *Store) UpdateStream(name string, s *stream.Stream) (err error) {
	// Get streamlist
	// if name not in streamList
	// 	abort tx, return error
	// serialize s to []byte
	// Put streamList
	// Put s to stream::name
	// Commit tx
	return
}

// OKAY

// initially let's just do
// streams stored at stream::<name>
// full list stored at stream_list
// For now we assume concurrent creation of streams is not a huge issue
// So:
// Creates pull list into memory, check for the name, append the new name
// Delete pull list into memory, check for the name, delete the existing name
// Lists pull the full list, grab the relevant names, retrieve those streams.

// NOW

// One missing ingredient: serialising streams.
