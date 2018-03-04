package badger

import "github.com/chulabs/seer/stream"

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
	// Get stream::name
	// if not found,
	// 	return error
	// deserialize into stream
	// return stream, nil
	return
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

// Problem here will be concurrent access on streamlist, which may also present issues
// for listing

// Really we want a list of names that we can quickly check for membership in and that
// we can get ordered segments of (i.e. offset limit).

// Okay we could have a hash like names::stream::foo which is either empty or not, that
// is quick and satisfies create and update.

// That doesn't fix paging though. For that we need some sorting mechanism. So redis
// ZRANGE is O(log(N)+M), which means it finds the start within log(N), then pages through
// a linked list.

// So I guess a zset would be implemented with a hash of members to structs of scores and pointers.

// Then if we add an element, we need to find its adjacent elements, set pointers to them, and update
// their pointers.

// So in both cases, how do we find an element given a score in log(N)?
