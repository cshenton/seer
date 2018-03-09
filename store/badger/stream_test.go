package badger_test

import (
	"testing"

	"github.com/chulabs/seer/store/badger"
	"github.com/chulabs/seer/stream"
)

func setUp(t *testing.T) (b *badger.Store) {
	b, err := badger.New()
	if err != nil {
		t.Fatal("unexpected error in badger.New:", err)
	}
	names := []string{"sales", "energy", "revenue"}

	for _, n := range names {
		s, err := stream.New(n, 3600, 0, 0, 0)
		if err != nil {
			t.Fatal("unexpected error in stream.New:", err)
		}
		err = b.CreateStream(n, s)
		if err != nil {
			t.Fatal("unexpected error in CreateStream:", err)
		}
	}

	return b
}

func TestCreateStream(t *testing.T) {

}

func TestCreateStreamErrs(t *testing.T) {

}

func TestGetStream(t *testing.T) {

}

func TestDeleteStream(t *testing.T) {

}

func TestUpdateStream(t *testing.T) {

}

func TestListStreams(t *testing.T) {

}
