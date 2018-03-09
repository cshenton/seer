package badger

import (
	"fmt"
	"log"

	// Avoid name conflict
	bdg "github.com/dgraph-io/badger"

	"github.com/chulabs/seer/store"
	"github.com/chulabs/seer/stream"
	"github.com/vmihailenco/msgpack"
)

// Key for the stream index
var streamListKey = []byte("stream_list")

// Key for a particular stream's data
func streamKey(name string) []byte {
	return []byte("stream::" + name)
}

// getStreamList retrieves the stream list using the provided, active transaction.
func getStreamList(tx *bdg.Txn) (l StreamList, err error) {
	item, err := tx.Get(streamListKey)
	if err != nil {
		return nil, &store.NoneFoundError{Kind: "streamList"}
	}
	lb, err := item.Value()
	if err != nil {
		log.Print(string(item.Key()))
		return nil, &store.CorruptDataError{Kind: "streamList"}
	}
	err = msgpack.Unmarshal(lb, &l)
	if err != nil {
		log.Print(lb)
		return nil, &store.CorruptDataError{Kind: "streamList"}
	}

	return l, nil
}

// streamInit initialises necessary keys to store streams.
func (b *Store) streamInit() {
	tx := b.NewTransaction(true)
	defer tx.Discard()

	_, err := getStreamList(tx)
	if err != nil {
		l := StreamList{"test"}
		lb, _ := msgpack.Marshal(l)
		tx.Set(streamListKey, lb)
		tx.Commit(nil)
	}
}

// CreateStream saves the provided stream at name, returns an error if a
// stream already exists at that address.
func (b *Store) CreateStream(name string, s *stream.Stream) (err error) {
	tx := b.NewTransaction(true)
	defer tx.Discard()

	l, err := getStreamList(tx)
	if err != nil {
		return err
	}
	if l.Contains(name) {
		return &store.AlreadyExistsError{Kind: "stream", Entity: name}
	}
	l = l.Add(name)

	sb, _ := msgpack.Marshal(s)
	tx.Set(streamKey(name), sb)
	lb, _ := msgpack.Marshal(l)
	tx.Set(streamListKey, lb)

	err = tx.Commit(nil)
	if err != nil {
		// conflict error
	}
	return nil
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

	l, err := getStreamList(tx)
	if err != nil {
		return err
	}
	if !l.Contains(name) {
		return &store.NotFoundError{Kind: "stream", Entity: name}
	}

	l = l.Remove(name)
	lb, _ := msgpack.Marshal(l)
	tx.Set(streamListKey, lb)

	tx.Delete(streamKey(name))
	err = tx.Commit(nil)
	if err != nil {
		// conflict error
	}
	return nil
}

// ListStreams returns a paged list of streams, or an error if none are found.
func (b *Store) ListStreams(pageNum, pageSize int) (s []*stream.Stream, err error) {
	tx := b.NewTransaction(true)
	defer tx.Discard()

	l, err := getStreamList(tx)
	if err != nil {
		return nil, err
	}
	fmt.Println(l)

	return
}

// UpdateStream overwrites the stream at name with the provided stream, or
// returns an error if no stream exists at name.
func (b *Store) UpdateStream(name string, s *stream.Stream) (err error) {
	tx := b.NewTransaction(true)
	defer tx.Discard()

	l, err := getStreamList(tx)
	if err != nil {
		return err
	}
	if !l.Contains(name) {
		return &store.NotFoundError{Kind: "stream", Entity: name}
	}

	sb, _ := msgpack.Marshal(s)
	tx.Set(streamKey(name), sb)
	err = tx.Commit(nil)
	if err != nil {
		// conflict error
	}
	return nil
}
