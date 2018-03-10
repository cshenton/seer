package model_test

import (
	"testing"

	"github.com/cshenton/seer/model"
	"gonum.org/v1/gonum/mat"
)

func TestBlockDiag(t *testing.T) {
	tt := []struct {
		name   string
		mats   []*mat.Dense
		output *mat.Dense
	}{
		{
			"Single square",
			[]*mat.Dense{mat.NewDense(2, 2, []float64{1, 0, 1, 1})},
			mat.NewDense(2, 2, []float64{1, 0, 1, 1}),
		},
		{
			"Single rectangle",
			[]*mat.Dense{mat.NewDense(3, 1, []float64{1, 2, 3})},
			mat.NewDense(3, 1, []float64{1, 2, 3}),
		},
		{
			"1x1 to 2x2",
			[]*mat.Dense{mat.NewDense(1, 1, []float64{1}), mat.NewDense(1, 1, []float64{1})},
			mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
		},
		{
			"1, 3 square to 4x4",
			[]*mat.Dense{mat.NewDense(1, 1, []float64{2}), mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})},
			mat.NewDense(4, 4, []float64{2, 0, 0, 0, 0, 1, 2, 3, 0, 4, 5, 6, 0, 7, 8, 9}),
		},
		{
			"1x2, 2x3 to 3x5",
			[]*mat.Dense{mat.NewDense(1, 2, []float64{2, 1}), mat.NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6})},
			mat.NewDense(3, 5, []float64{2, 1, 0, 0, 0, 0, 0, 1, 2, 3, 0, 0, 4, 5, 6}),
		},
		{
			"Four 1x1 to 4x4",
			[]*mat.Dense{mat.NewDense(1, 1, []float64{1}), mat.NewDense(1, 1, []float64{2}), mat.NewDense(1, 1, []float64{3}), mat.NewDense(1, 1, []float64{4})},
			mat.NewDense(4, 4, []float64{1, 0, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 4}),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b := model.BlockDiag(tc.mats...)
			if !mat.Equal(b, tc.output) {
				t.Errorf("Expected result %v, but got %v", tc.output, b)
			}
		})
	}
}

func TestDenseValues(t *testing.T) {
	tt := []struct {
		Name   string
		RowDim int
		ColDim int
		Data   []float64
	}{
		{"Single element", 1, 1, []float64{1.0}},
		{"Simple 2x2", 2, 2, []float64{1, 2, 3, 4}},
		{"Simple 3x3", 3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"Decimals 3x4", 3, 4, []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 71.3, 0.0, 0.0, 0.0, 1.21, 0.04}},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			m := mat.NewDense(tc.RowDim, tc.ColDim, tc.Data)
			values := model.DenseValues(m)
			for i, v := range values {
				if v != tc.Data[i] {
					t.Fatalf("output value %v did not match input %v at position %v", v, tc.Data[i], i)
					return
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

func TestEye(t *testing.T) {
	tt := []struct {
		name string
		dim  int
		mat  *mat.Dense
	}{
		{"Dimension 1", 1, mat.NewDense(1, 1, []float64{1})},
		{"Dimension 2", 2, mat.NewDense(2, 2, []float64{1, 0, 0, 1})},
		{"Dimension 3", 3, mat.NewDense(3, 3, []float64{1, 0, 0, 0, 1, 0, 0, 0, 1})},
		{"Dimension 4", 4, mat.NewDense(4, 4, []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			m, err := model.Eye(tc.dim)
			if err != nil {
				t.Fatal("failed to create matrix", err)
			}
			if !mat.Equal(tc.mat, m) {
				t.Errorf("expected result %v, but got %v", tc.mat, m)
			}
		})
	}
}

func TestEyeErrs(t *testing.T) {
	_, err := model.Eye(0)

	if err == nil {
		t.Errorf("expected error, but it was nil")
	}
}
