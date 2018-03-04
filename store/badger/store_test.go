package badger_test

import (
	"testing"

	"github.com/chulabs/seer/store/badger"
)

func TestNew(t *testing.T) {
	_, err := badger.New()
	if err != nil {
		t.Error("unexpected error in New,", err)
	}
}
