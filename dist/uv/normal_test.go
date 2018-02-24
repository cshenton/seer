package uv_test

import (
	"math"
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
				t.Errorf("expected location %v, but got %v", tc.loc, n.Location)
			}
			if n.Scale != tc.scale {
				t.Errorf("expected Scale %v, but got %v", tc.scale, n.Scale)
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

func TestNormalMeanAndVariance(t *testing.T) {
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
			if n.Mean() != tc.loc {
				t.Errorf("expected mean %v, but got %v", tc.loc, n.Mean())
			}
			if n.Variance() != math.Pow(tc.scale, 2) {
				t.Errorf("expected variance %v, but got %v", math.Pow(tc.scale, 2), n.Variance())
			}
		})
	}
}

func TestNormalQuantile(t *testing.T) {
	loc := 0.0
	scale := 1.0

	n, err := uv.NewNormal(loc, scale)
	if err != nil {
		t.Error("unexpected error in NewNormal,", err)
	}
	if math.Abs(n.Quantile(0.5)-loc) > 1e-8 {
		t.Errorf("expected median quantile %v, but got %v", loc, n.Quantile(0.5))
	}
}
