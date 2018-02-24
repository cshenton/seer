package uv_test

import (
	"testing"

	"github.com/chulabs/seer/dist/uv"
)

func TestNewNormal(t *testing.T) {
	tt := []struct {
		name  string
		loc   float64
		scale float64
	}{
		{"standard", 0, 1},
		{"negative loc", -10, 2},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			n, err := uv.NewNormal(tc.loc, tc.scale)
			if err != nil {
				t.Error("unexpected error in NewNormal,", err)
			}
			if n.Location != tc.loc {
				t.Errorf("expected location %v, but got %v", n.Location, tc.loc)
			}
			if n.Scale != tc.scale {
				t.Errorf("expected Scale %v, but got %v", n.Scale, tc.scale)
			}

		})
	}
}

func TestNewNormalErrs(t *testing.T) {
	tt := []struct {
		name  string
		loc   float64
		scale float64
	}{
		{"zero scale", 10, 0},
		{"negative scale", 10, -2},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			n, err := uv.NewNormal(tc.loc, tc.scale)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
			if n != nil {
				t.Error("expected nil dist but it was", n)
			}
		})
	}
}
