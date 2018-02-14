package model_test

import (
	"testing"

	"github.com/chulabs/seer/model"
)

// yr, seconds in an average year (365.25 days)
const yr = float64(31577600)

func TestHarmonics(t *testing.T) {
	tt := []struct {
		name   string
		min    float64
		max    float64
		result []float64
	}{
		{
			"1 and 10", 1, 10, []float64{
				10.0, 10.0 / 2, 10.0 / 3, 10.0 / 4, 10.0 / 5, 10.0 / 6, 10.0 / 7, 10.0 / 8,
				10.0 / 9, 10.0 / 10.0,
			},
		},
		{
			"1 and 100", 1, 100, []float64{
				100.0, 100.0 / 2, 100.0 / 3, 100.0 / 4, 100.0 / 5, 100.0 / 6, 100.0 / 7, 100.0 / 8,
				100.0 / 9, 100.0 / 10,
				100.0 / 20, 100.0 / 30, 100.0 / 40, 100.0 / 50, 100.0 / 60, 100.0 / 70, 100.0 / 80,
				100.0 / 90, 100.0 / 100,
			},
		},
		{
			"Hourly", 3600, yr,
			[]float64{
				yr, yr / 2, yr / 3, yr / 4, yr / 5, yr / 6, yr / 7, yr / 8, yr / 9, yr / 10,
				yr / 20, yr / 30, yr / 40, yr / 50, yr / 60, yr / 70, yr / 80, yr / 90, yr / 100,
				yr / 200, yr / 300, yr / 400, yr / 500, yr / 600, yr / 700, yr / 800, yr / 900, yr / 1000,
			},
		},
		{
			"Daily", 86400, yr,
			[]float64{
				yr, yr / 2, yr / 3, yr / 4, yr / 5, yr / 6, yr / 7, yr / 8, yr / 9, yr / 10,
				yr / 20, yr / 30, yr / 40, yr / 50, yr / 60, yr / 70, yr / 80, yr / 90, yr / 100,
			},
		},
		{
			"Weekly", 604800, yr,
			[]float64{
				yr, yr / 2, yr / 3, yr / 4, yr / 5, yr / 6, yr / 7, yr / 8, yr / 9, yr / 10,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			h := model.Harmonics(tc.min, tc.max)
			if len(h) != len(tc.result) {
				t.Fatalf("different lengths, expected: %v, got: %v", len(tc.result), len(h))
			}

			for i := range tc.result {
				if h[i] != tc.result[i] {
					t.Fatalf("expected %v at position %v, but got %v", tc.result[i], i, h[i])
				}
			}
		})
	}
}

func TestDiag(t *testing.T) {
	tt := []struct {
		name string
		v    []float64
		d    []float64
	}{
		{"1x1", []float64{1}, []float64{1}},
		{"2x2", []float64{2, 1}, []float64{2, 0, 0, 1}},
		{"3x3", []float64{3, 2, 1}, []float64{3, 0, 0, 0, 2, 0, 0, 0, 1}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			d := model.Diag(tc.v)
			if len(d) != len(tc.d) {
				t.Fatalf("expected result length %v, got %v", len(tc.d), len(d))
			}
			for i := range d {
				if d[i] != tc.d[i] {
					t.Errorf("mismatch at %v, expected %v, got %v", i, tc.d[i], d[i])
				}
			}
		})
	}
}

func TestNewDeterministic(t *testing.T) {
	loc := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	cov := model.Diag([]float64{
		1e15, 1e5, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4,
		1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4,
	})
	d := model.NewDeterministic(604800)

	if len(loc) != len(d.Location) {
		t.Fatalf("location length was %v, expected %v", len(d.Location), len(loc))
	}
	if len(cov) != len(d.Covariance) {
		t.Fatalf("covariance length was %v, expected %v", len(d.Covariance), len(cov))
	}

	for i := range loc {
		if loc[i] != d.Location[i] {
			t.Errorf("location expected %v at position %v, but got %v", loc[i], i, d.Location[i])
		}
	}
	for i := range cov {
		if cov[i] != d.Covariance[i] {
			t.Errorf("covariance expected %v at position %v, but got %v", cov[i], i, d.Covariance[i])
		}
	}
}
