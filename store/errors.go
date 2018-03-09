package store

import (
	"fmt"
)

// NotFoundError is returned when the requested Entity is not in the data store.
type NotFoundError struct {
	Kind   string
	Entity string
}

// Error message for NotFoundErrors, implements error interface.
func (err *NotFoundError) Error() string {
	return fmt.Sprintf("no %v with name %v was found in store", err.Kind, err.Entity)
}

// AlreadyExistsError is returned when an Entity name is already taken.
type AlreadyExistsError struct {
	Kind   string
	Entity string
}

// Error message for AlreadyExistsError, implements error interface.
func (err *AlreadyExistsError) Error() string {
	return fmt.Sprintf("a %v with name %v already exists", err.Kind, err.Entity)
}

// NoneFoundError is returned when a list request on a Kind returns no entities.
type NoneFoundError struct {
	Kind string
}

// Error message for NoneFoundError, implements error interface.
func (err *NoneFoundError) Error() string {
	return fmt.Sprintf("no entities of kind %v were found", err.Kind)
}

// CorruptDataError is returned when data at a key has invalid schema.
type CorruptDataError struct {
	Kind string
}

// Error message for CorruptDataError, implements error interface.
func (err *CorruptDataError) Error() string {
	return fmt.Sprintf("unable to unmarshal entity of kind %v", err.Kind)
}
