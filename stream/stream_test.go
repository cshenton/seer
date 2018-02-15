package stream_test

import (
	"testing"

	"github.com/chulabs/seer/stream"
)

func TestNew(t *testing.T) {
	s, err := stream.New("sales", 60, 10, 0, 1)

	if err != nil {
		t.Error("unexpected error in New:", err)
	}
	if s.Config.Min != 10 {
		t.Errorf("expected config min %v, but got %v", 10, s.Config.Min)
	}
}

func TestNewErrs(t *testing.T) {
	s, err := stream.New("y", 60, 0, 0, 0)

	if err == nil {
		t.Error("expected error but got nil")
	}
	if s != nil {
		t.Error("expected nil stream but got:", s)
	}
}
