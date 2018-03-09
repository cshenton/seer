package bolt_test

import (
	"testing"

	"github.com/chulabs/seer/store/bolt"
	"github.com/chulabs/seer/stream"
)

func setUp(t *testing.T) (b *bolt.Store) {
	b, err := bolt.New(testPath(t))
	if err != nil {
		t.Fatal("unexpected error in bolt.New:", err)
	}
	names := []string{"sales", "visits", "usage"}

	for _, n := range names {
		s, _ := stream.New(n, 3600, 0, 0, 0)
		err := b.CreateStream(n, s)
		if err != nil {
			t.Fatal("unexpected error in CreateStream:", err)
		}
	}

	return b
}

func TestCreateStreamErrs(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	s, _ := stream.New("sales", 3600, 0, 0, 0)
	err := b.CreateStream("sales", s)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}

func TestGetStream(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	tt := []string{"sales", "visits", "usage"}

	for _, name := range tt {
		t.Run(name, func(t *testing.T) {
			s, err := b.GetStream(name)
			if err != nil {
				t.Error("unexpected error in GetStream:", err)
			}
			if s.Config.Name != name {
				t.Errorf("expected stream name %v, but got %v", name, s.Config.Name)
			}
		})
	}
}

func TestGetStreamErrs(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	s, err := b.GetStream("notastream")
	if err == nil {
		t.Error("expected error, but it was nil")
	}
	if s != nil {
		t.Error("expected nil stream, but it was", s)
	}
}

func TestDeleteStream(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	tt := []string{"sales", "visits", "usage"}

	for _, name := range tt {
		t.Run(name, func(t *testing.T) {
			err := b.DeleteStream(name)
			if err != nil {
				t.Error("unexpected error in DeleteStream:", err)
			}

			_, err = b.GetStream(name)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
		})
	}
}

func TestDeleteStreamErrs(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	err := b.DeleteStream("notastream")
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}
