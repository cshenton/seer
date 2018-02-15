package kalman_test

import (
	"testing"

	"github.com/chulabs/seer/kalman"
	"gonum.org/v1/gonum/mat"
)

func TestNewState(t *testing.T) {
	tt := []struct {
		name   string
		loc    *mat.Dense
		cov    *mat.Dense
		errNil bool
	}{
		{"Non column loc", mat.NewDense(1, 2, []float64{0, 0}), mat.NewDense(1, 1, []float64{1}), false},
		{"Non square cov", mat.NewDense(1, 1, []float64{0}), mat.NewDense(1, 2, []float64{1, 1}), false},
		{"Non matching loc, cov", mat.NewDense(2, 1, []float64{0, 0}), mat.NewDense(1, 1, []float64{1}), false},
		{"Valid dim 1", mat.NewDense(1, 1, []float64{0}), mat.NewDense(1, 1, []float64{1}), true},
		{"Valid dim 2", mat.NewDense(2, 1, []float64{0, 0}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}), true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := kalman.NewState(tc.loc, tc.cov)
			errNil := (err == nil)
			if errNil != tc.errNil {
				t.Errorf("Expected error == nil to be %v, but it was %v", tc.errNil, errNil)
			}
		})
	}
}
