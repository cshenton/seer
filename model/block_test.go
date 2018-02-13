package model_test

import (
	"testing"

	"github.com/chulabs/seer/model"
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
