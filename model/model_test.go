package model_test

import (
	"testing"

	"github.com/chulabs/seer/model"
)

func TestNew(t *testing.T) {
	m := model.New(604800)

	if len(m.Deterministic.Location) != 22 {
		t.Errorf("expected deterministic location dim of %v, but got %v", 22, len(m.Deterministic.Location))
	}
	if len(m.Stochastic.Location) != 1 {
		t.Errorf("expected stochastic location dim of %v, but got %v", 1, len(m.Stochastic.Location))
	}
	if m.RCE.Theta.Scale != 180 {
		t.Errorf("expected initial theta scale of %v, but got %v", 180, m.RCE.Theta.Scale)
	}
}
