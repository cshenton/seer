package uv_test

import (
	"testing"

	"github.com/cshenton/seer/dist/uv"
)

func TestNewInverseGamma(t *testing.T) {
	tt := []struct {
		name  string
		shape float64
		scale float64
	}{
		{"Small", 0.5, 0.6},
		{"Large", 30, 25},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			i, err := uv.NewInverseGamma(tc.shape, tc.scale)
			if err != nil {
				t.Error("unexpected error:", err)
			}
			if i.Shape != tc.shape {
				t.Errorf("expected shape %v, but got %v", tc.shape, i.Shape)
			}
		})
	}
}

func TestNewInverseGammaErrs(t *testing.T) {
	tt := []struct {
		name  string
		shape float64
		scale float64
	}{
		{"Negative Shape", -10, 5},
		{"Zero Scale", 10, 0},
		{"Both Negative", -1, -1},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			i, err := uv.NewInverseGamma(tc.shape, tc.scale)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
			if i != nil {
				t.Error("expected nil pointer, but it was", i)
			}
		})
	}
}

func TestInverseGammaMean(t *testing.T) {
	tt := []struct {
		name  string
		shape float64
		scale float64
		mean  float64
	}{
		{"Medium", 4, 5, 5.0 / 3.0},
		{"Large", 30, 25, 25.0 / 29.0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			i, err := uv.NewInverseGamma(tc.shape, tc.scale)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}
			if i.Mean() != tc.mean {
				t.Errorf("expected mean %v, but it was %v", tc.mean, i.Mean())
			}
		})
	}
}

func TestInverseGammaVariance(t *testing.T) {
	tt := []struct {
		name     string
		shape    float64
		scale    float64
		variance float64
	}{
		{"Medium", 4, 5, 25.0 / 9.0 / 2.0},
		{"Large", 30, 25, 625.0 / 841.0 / 28.0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			i, err := uv.NewInverseGamma(tc.shape, tc.scale)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}
			if i.Variance() != tc.variance {
				t.Errorf("expected mean %v, but it was %v", tc.variance, i.Variance())
			}
		})
	}
}
