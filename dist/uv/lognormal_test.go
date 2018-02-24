package uv_test

import (
	"testing"

	"github.com/chulabs/seer/dist/uv"
)

func TestNewLogNormal(t *testing.T) {
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
			ln, err := uv.NewLogNormal(tc.loc, tc.scale)
			if err != nil {
				t.Error("unexpected error in NewLogNormal,", err)
			}
			if ln.Location != tc.loc {
				t.Errorf("expected location %v, but got %v", ln.Location, tc.loc)
			}
			if ln.Scale != tc.scale {
				t.Errorf("expected Scale %v, but got %v", ln.Scale, tc.scale)
			}

		})
	}
}

func TestNewLogNormalErrs(t *testing.T) {
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
			ln, err := uv.NewLogNormal(tc.loc, tc.scale)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
			if ln != nil {
				t.Error("expected nil dist but it was", ln)
			}
		})
	}
}
