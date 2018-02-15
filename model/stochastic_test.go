package model_test

import (
	"testing"

	"github.com/chulabs/seer/model"
)

func TestNewStochastic(t *testing.T) {
	s := model.NewStochastic()

	if s.Location[0] != 0.0 {
		t.Error("Expected location 0.0 but got", s.Location[0])
	}
	if s.Covariance[0] != 1e12 {
		t.Error("Expected covariance 1e12, but got", s.Covariance[0])
	}
}

func TestStochasticSystem(t *testing.T) {
	s := model.NewStochastic()
	k := s.System(100, 10)

	if k.Q.At(0, 0) != 10 {
		t.Errorf("expected process covariance of %v, but it was %v", 10, k.Q.At(0, 0))
	}
	if k.R.At(0, 0) != 100 {
		t.Errorf("expected observation covariance of %v, but it was %v", 100, k.R.At(0, 0))
	}
}
