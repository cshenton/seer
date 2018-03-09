package store

import "context"

// StreamMiddleware appends the provided StreamStore to the request context.
func StreamMiddleware(c context.Context, s StreamStore) context.Context {
	return streamToContext(c, s)
}
