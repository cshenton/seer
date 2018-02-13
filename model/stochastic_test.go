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
