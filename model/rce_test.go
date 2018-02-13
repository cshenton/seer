package model_test

import (
	"testing"

	"github.com/chulabs/seer/model"
)

func TestHistoryUpdate(t *testing.T) {
	h := model.History([]float64{1, 2})
	h.Update(3)

	if h[0] != 3 {
		t.Errorf("expected entry 0 to be %v, but it was %v", 3, h[0])
	}
	if h[1] != 1 {
		t.Errorf("expected entry 1 to be %v, but it was %v", 1, h[1])
	}
}
