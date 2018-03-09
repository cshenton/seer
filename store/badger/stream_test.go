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

func tearDown(b *badger.Store) {
	names := []string{"sales", "energy", "revenue"}

	for _, n := range names {
		b.DeleteStream(n)
	}

	b.Close()
}

func TestCreateStreamErrs(t *testing.T) {
	b := setUp(t)
	defer tearDown(b)

	name := "sales"

	s, err := stream.New(name, 3600, 0, 0, 0)
	if err != nil {
		t.Fatal("unexpected error in stream.New:", err)
	}
	err = b.CreateStream(name, s)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}

func TestGetStream(t *testing.T) {
	b := setUp(t)
	defer tearDown(b)

	tt := []struct {
		name string
	}{
		{"sales"},
		{"energy"},
		{"revenue"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := b.GetStream(tc.name)
			if err != nil {
				t.Error("unexpected error in GetStream:", err)
			}
			if s.Config.Name != tc.name {
				t.Errorf("expected config name %v, but got %v", tc.name, s.Config.Name)
			}
		})
	}
}

func TestDeleteStreamErrs(t *testing.T) {
	b := setUp(t)
	defer tearDown(b)

	err := b.DeleteStream("notastream")
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}

func TestUpdateStream(t *testing.T) {

}

func TestListStreams(t *testing.T) {

}
