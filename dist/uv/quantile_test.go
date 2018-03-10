package uv_test

import (
	"math"
	"testing"

	"github.com/cshenton/seer/dist/uv"
)

func TestConfidenceInterval(t *testing.T) {
	n, _ := uv.NewNormal(0, 1)
	l, u, err := uv.ConfidenceInterval(n, 0.9)
	if err != nil {
		t.Fatal("unexpected error in ConfidenceInterval,", err)
	}
	if math.Abs(l - -1.645) > 1e-3 {
		t.Errorf("expected lower bound %v, but got %v", -1.645, l)
	}
	if math.Abs(u-1.645) > 1e-3 {
		t.Errorf("expected upper bound %v, but got %v", 1.645, u)
	}

	_, _, err = uv.ConfidenceInterval(n, -1)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}
