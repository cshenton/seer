package store

import (
	"context"

	"github.com/cshenton/seer/stream"
)

// StreamStore defines the methods required to store and retrieve streams.
type StreamStore interface {
	CreateStream(name string, s *stream.Stream) (err error)
	GetStream(name string) (s *stream.Stream, err error)
	DeleteStream(name string) (err error)
	ListStreams(pageNum, pageSize int) (s []*stream.Stream, err error)
	UpdateStream(name string, s *stream.Stream) (err error)
}

// CreateStream creates a stream using the store on the current context, it returns an
// error if the stream already exists.
func CreateStream(c context.Context, name string, s *stream.Stream) (err error) {
	return streamFromContext(c).CreateStream(name, s)
}

// GetStream returns the stream with the specific name using the current context store.
func GetStream(c context.Context, name string) (s *stream.Stream, err error) {
	return streamFromContext(c).GetStream(name)
}

// DeleteStream deletes the stream with the specific name using the current context store.
func DeleteStream(c context.Context, name string) (err error) {
	return streamFromContext(c).DeleteStream(name)
}

// ListStreams lists pages streams from the current context store.
func ListStreams(c context.Context, pageNum, pageSize int) (s []*stream.Stream, err error) {
	return streamFromContext(c).ListStreams(pageNum, pageSize)
}

// UpdateStream saves the provided stream, and returns an error if no stream with
// the given name exists.
func UpdateStream(c context.Context, name string, s *stream.Stream) (err error) {
	return streamFromContext(c).UpdateStream(name, s)
}
