package store

import "context"

type contextKey string

var contextKeyStream = contextKey("stream")

// FromContext returns the StreamStore associated with this context.
func streamFromContext(c context.Context) StreamStore {
	return c.Value(contextKeyStream).(StreamStore)
}

// ToContext returns a new context with value StreamStore.
func streamToContext(c context.Context, s StreamStore) context.Context {
	return context.WithValue(c, contextKeyStream, s)
}
