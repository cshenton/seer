package store

import "context"

// BadgerMiddleware appends the badger implementation of the store interfaces
// onto the current request context.
func BadgerMiddleware(c context.Context) context.Context {
	type tempString string
	return context.WithValue(c, tempString("foo"), "bar")
	// do this once interface implemented
	// return streamToContext(c, s)
}
