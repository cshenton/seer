package model_test

import (
	"testing"

	"github.com/cshenton/seer/model"
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

func TestModelUpdate(t *testing.T) {
	m := model.New(604800)

	m.Update(604800, 1.0)

	if m.Deterministic.Location[0] == 0 {
		t.Error("deterministic not updated")
	}

	if m.Stochastic.Location[0] == 0 {
		t.Error("stochastic not updated")
	}

	if m.RCE.History[0] == 0 {
		t.Error("RCE not updated")
	}
}

func TestModelForecast(t *testing.T) {
	period := 604800.0
	n := 150

	m := model.New(period)
	m.Update(period, 1)
	f := m.Forecast(period, n)

	if len(f) != n {
		t.Errorf("expected length %v, but it was %v", n, len(f))
	}
}
