package badger

import (
	"fmt"
	"log"

	"github.com/chulabs/seer/store"
	"github.com/chulabs/seer/stream"
	"github.com/vmihailenco/msgpack"
)

type streamList []string

// Key for the stream index
var streamListKey = []byte("stream_list")

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

// GetStream returns the stream stored at name, or an error if the stream does
// not exist, or has corrupted data.
func (b *Store) GetStream(name string) (s *stream.Stream, err error) {
	s = &stream.Stream{}
	tx := b.NewTransaction(false)
	defer tx.Discard()

	item, err := tx.Get(streamKey(name))
	if err != nil {
		return nil, &store.NotFoundError{Kind: "stream", Entity: name}
	}
	sb, err := item.Value()
	if err != nil {
		log.Print(string(item.Key()))
		return nil, &store.CorruptDataError{Kind: "stream"}
	}
	err = msgpack.Unmarshal(sb, s)
	if err != nil {
		log.Print(sb)
		return nil, &store.CorruptDataError{Kind: "stream"}
	}

	tx.Commit(nil)
	return s, nil
}

// DeleteStream deletes the stream stored at name, or returns an error if
// no such stream exists.
func (b *Store) DeleteStream(name string) (err error) {
	tx := b.NewTransaction(true)
	defer tx.Discard()

	err = tx.Delete(streamKey(name))
	if err != nil {
		return &store.NotFoundError{Kind: "stream", Entity: name}
	}

	list, err := tx.Get(streamListKey)
	if err != nil {
		return &store.NoneFoundError{Kind: "streamList"}
	}
	// delete stream from list
	fmt.Println(list)
	// delete stream from list

	err = tx.Commit(nil)
	if err != nil {
		// conflict error
	}

	return nil
}

// ListStreams returns a paged list of streams, or an error if none are found.
func (b *Store) ListStreams(pageNum, pageSize int) (s []*stream.Stream, err error) {
	return
}

// UpdateStream overwrites the stream at name with the provided stream, or
// returns an error if no stream exists at name.
func (b *Store) UpdateStream(name string, s *stream.Stream) (err error) {
	var l streamList
	tx := b.NewTransaction(true)
	defer tx.Discard()

	item, err := tx.Get(streamListKey)
	if err != nil {
		return &store.NoneFoundError{Kind: "streamList"}
	}
	lb, err := item.Value()
	if err != nil {
		log.Print(string(item.Key()))
		return &store.CorruptDataError{Kind: "streamList"}
	}
	err = msgpack.Unmarshal(lb, l)
	if err != nil {
		log.Print(lb)
		return &store.CorruptDataError{Kind: "streamList"}
	}
	// if name not in streamList
	// 	abort tx, return error

	sb, _ := msgpack.Marshal(s)
	tx.Set(streamKey(name), sb)
	err = tx.Commit(nil)
	if err != nil {
		// conflict error
	}
	return nil
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
