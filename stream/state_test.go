package stream_test

import (
	"testing"

	"github.com/chulabs/seer/stream"
)

func TestNewState(t *testing.T) {
	c, err := stream.NewConfig("test", 604800, 0, 0, 0)
	if err != nil {
		t.Fatal("unexpected error in NewConfig:", err)
	}
	s := stream.NewState(c)

	if len(s.Deterministic.Location) != 22 {
		t.Errorf("expected deterministic location dim of %v, but got %v", 22, len(s.Deterministic.Location))
	}
	if len(s.Stochastic.Location) != 1 {
		t.Errorf("expected stochastic location dim of %v, but got %v", 1, len(s.Stochastic.Location))
	}
	if s.RCE.Theta.Scale != 180 {
		t.Errorf("expected initial theta scale of %v, but got %v", 180, s.RCE.Theta.Scale)
	}
}
