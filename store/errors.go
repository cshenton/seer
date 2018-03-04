package store

import (
	"fmt"
)

// NotFoundError is returned when the requested entity is not in the data store.
type NotFoundError struct {
	kind   string
	entity string
}

// Error message for NotFoundErrors, implements error interface.
func (err *NotFoundError) Error() string {
	return fmt.Sprintf("no %v with name %v was found in store", err.kind, err.entity)
}

// AlreadyExistsError is returned when an entity name is already taken.
type AlreadyExistsError struct {
	kind   string
	entity string
}

// Error message for AlreadyExistsError, implements error interface.
func (err *AlreadyExistsError) Error() string {
	return fmt.Sprintf("a %v with name %v already exists", err.kind, err.entity)
}

// NoneFoundError is returned when a list request on a kind returns no entities.
type NoneFoundError struct {
	kind string
}

// Error message for NoneFoundError, implements error interface.
func (err *NoneFoundError) Error() string {
	return fmt.Sprintf("no entities of kind %v were found", err.kind)
}
