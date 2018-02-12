package store

import "context"

// StreamStore defines the methods required to store and retrieve streams.
type StreamStore interface {
	CreateStream()
	GetStream()
	DeleteStream()
	ListStreams()
	UpdateStream()
}

// CreateStream creates a stream using the store on the current context.
func CreateStream(c context.Context) {
	streamFromContext(c).CreateStream()
}
