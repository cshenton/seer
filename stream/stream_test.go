package stream_test

import (
	"testing"
	"time"

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

func TestStreamUpdate(t *testing.T) {
	vals := []float64{1, 2}
	times := []time.Time{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC)}

	s, err := stream.New("streamy", 86400, 0, 0, 0)
	if err != nil {
		t.Fatal("unexpected error in New:", err)
	}

	err = s.Update(vals, times)
	if err != nil {
		t.Fatal("unexpected error in Update:", err)
	}

	if s.Model.RCE.History[0] == 0 {
		t.Error("model was not updated")
	}
}

func TestStreamUpdateErrs(t *testing.T) {

}
