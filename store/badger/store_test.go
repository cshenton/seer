package badger_test

import (
	"testing"

	"github.com/chulabs/seer/store/badger"
)

func TestNew(t *testing.T) {
	b, err := badger.New()
	defer b.Close()

	if err != nil {
		t.Error("unexpected error in New,", err)
	}
}
