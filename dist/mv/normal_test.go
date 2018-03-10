package mv_test

import (
	"testing"

	"github.com/cshenton/seer/dist/mv"
)

func TestNewNormal(t *testing.T) {
	tt := []struct {
		name string
		loc  []float64
		cov  []float64
	}{
		{"1x1", []float64{1}, []float64{1}},
		{"2x2", []float64{1, 1}, []float64{1, 0, 0, 1}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			n, err := mv.NewNormal(tc.loc, tc.cov)
			if err != nil {
				t.Error("unexpected error in NewNormal:", err)
			}
			if n.Location[0] != tc.loc[0] {
				t.Error("non matching input, output loc")
			}
		})
	}
}

func TestNewNormalErrs(t *testing.T) {
	tt := []struct {
		name string
		loc  []float64
		cov  []float64
	}{
		{"empty location", []float64{}, []float64{1}},
		{"mismatched lengths", []float64{1}, []float64{1, 1}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			n, err := mv.NewNormal(tc.loc, tc.cov)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
			if n != nil {
				t.Error("expected nil pointer, but it was", n)
			}
		})
	}
}

func TestDim(t *testing.T) {
	tt := []struct {
		name string
		loc  []float64
		cov  []float64
		dim  int
	}{
		{"1x1", []float64{1}, []float64{1}, 1},
		{"2x2", []float64{1, 1}, []float64{1, 0, 0, 1}, 2},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			n, err := mv.NewNormal(tc.loc, tc.cov)
			if err != nil {
				t.Fatal(err)
			}
			if tc.dim != n.Dim() {
				t.Errorf("expected Dim %v, but it was %v", tc.dim, n.Dim())
			}
		})
	}
}
