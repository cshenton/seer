package store_test

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/chulabs/seer/store"
	"github.com/chulabs/seer/store/bolt"
	"github.com/chulabs/seer/stream"
)

func testPath(t *testing.T) string {
	f, err := ioutil.TempFile(os.TempDir(), "bolt_test")
	if err != nil {
		t.Fatal("failed to create test db file")
	}
	return f.Name()
}

func setUp(t *testing.T) (c context.Context) {
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

	c = context.Background()
	c = store.StreamMiddleware(c, b)
	return c
}

func TestCreateStream(t *testing.T) {
	c := setUp(t)
	name := "test"

	s, _ := stream.New(name, 3600, 0, 0, 0)

	err := store.CreateStream(c, name, s)
	if err != nil {
		t.Error("unexpected error in CreateStream:", err)
	}
}

func TestGetStream(t *testing.T) {
	c := setUp(t)
	name := "sales"

	s, err := store.GetStream(c, name)
	if err != nil {
		t.Error("unexpected error in GetStream:", err)
	}
	if s.Config.Name != name {
		t.Errorf("expected name %v, but got %v", name, s.Config.Name)
	}
}

func TestDeleteStream(t *testing.T) {
	c := setUp(t)
	name := "sales"

	err := store.DeleteStream(c, name)
	if err != nil {
		t.Error("unexpected error in DeleteStream:", err)
	}
}

func TestUpdateStream(t *testing.T) {
	c := setUp(t)
	name := "sales"

	s, _ := stream.New(name, 3600, 0, 0, 0)

	err := store.UpdateStream(c, name, s)
	if err != nil {
		t.Error("unexpected error in UpdateStream:", err)
	}
}

func TestListStreams(t *testing.T) {
	c := setUp(t)

	s, err := store.ListStreams(c, 1, 2)
	if err != nil {
		t.Error("unexpected error in GetStream:", err)
	}
	if len(s) != 2 {
		t.Errorf("expected %v streams, but got %v", 2, len(s))
	}
}
