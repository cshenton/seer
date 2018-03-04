package store_test

import (
	"testing"

	"github.com/chulabs/seer/store"
)

func TestNotFoundError(t *testing.T) {
	msg := "no stream with name wallace was found in store"
	err := store.NotFoundError{
		Kind:   "stream",
		Entity: "wallace",
	}

	if err.Error() != msg {
		t.Errorf("expected message `%v`, but got `%v`", msg, err.Error())
	}
}

func TestAlreadyExistsError(t *testing.T) {
	msg := "a stream with name wallace already exists"
	err := store.AlreadyExistsError{
		Kind:   "stream",
		Entity: "wallace",
	}

	if err.Error() != msg {
		t.Errorf("expected message `%v`, but got `%v`", msg, err.Error())
	}
}

func TestNoneFoundError(t *testing.T) {
	msg := "no entities of kind stream were found"
	err := store.NoneFoundError{
		Kind: "stream",
	}

	if err.Error() != msg {
		t.Errorf("expected message `%v`, but got `%v`", msg, err.Error())
	}
}

func TestCorruptDataError(t *testing.T) {
	msg := "unable to unmarshal entity of kind stream"
	err := store.CorruptDataError{
		Kind: "stream",
	}

	if err.Error() != msg {
		t.Errorf("expected message `%v`, but got `%v`", msg, err.Error())
	}
}
