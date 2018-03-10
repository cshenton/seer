/*
 * Copyright (C) 2018 The Seer Authors. All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package kalman_test

import (
	"testing"

	"github.com/cshenton/seer/kalman"
	"gonum.org/v1/gonum/mat"
)

func TestNewSystem(t *testing.T) {
	tt := []struct {
		name   string
		a      *mat.Dense
		b      *mat.Dense
		c      *mat.Dense
		q      *mat.Dense
		r      *mat.Dense
		errNil bool
	}{
		{
			"Q not square", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			mat.NewDense(1, 2, []float64{1, 0}), mat.NewDense(2, 1, []float64{.5, 0}), mat.NewDense(1, 1, []float64{.5}), false,
		},
		{
			"R not square", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			mat.NewDense(1, 2, []float64{1, 0}), mat.NewDense(2, 2, []float64{.5, 0, 0, .5}), mat.NewDense(1, 2, []float64{0, .5}), false,
		},
		{
			"A not square", mat.NewDense(3, 1, []float64{1, 1, 0}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			mat.NewDense(1, 2, []float64{1, 0}), mat.NewDense(2, 2, []float64{.5, 0, 0, .5}), mat.NewDense(1, 1, []float64{.5}), false,
		},
		{
			"A, B don't match", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(3, 1, []float64{1, 0, 0}),
			mat.NewDense(1, 2, []float64{1, 0}), mat.NewDense(2, 2, []float64{.5, 0, 0, .5}), mat.NewDense(1, 1, []float64{.5}), false,
		},
		{
			"B, Q don't match", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			mat.NewDense(1, 2, []float64{1, 0}), mat.NewDense(1, 1, []float64{.5}), mat.NewDense(1, 1, []float64{.5}), false,
		},
		{
			"C, R don't match", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			mat.NewDense(2, 1, []float64{1, 0}), mat.NewDense(2, 2, []float64{.5, 0, 0, .5}), mat.NewDense(1, 1, []float64{.5}), false,
		},
		{
			"C, A don't match", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			mat.NewDense(1, 1, []float64{1}), mat.NewDense(2, 2, []float64{.5, 0, 0, .5}), mat.NewDense(1, 1, []float64{.5}), false,
		},
		{
			"Valid with square B", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			mat.NewDense(1, 2, []float64{1, 0}), mat.NewDense(2, 2, []float64{.5, 0, 0, .5}), mat.NewDense(1, 1, []float64{.5}), true,
		},
		{
			"Valid with non square B", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(2, 1, []float64{1, 1}),
			mat.NewDense(1, 2, []float64{1, 0}), mat.NewDense(1, 1, []float64{.5}), mat.NewDense(1, 1, []float64{.5}), true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := kalman.NewSystem(tc.a, tc.b, tc.c, tc.q, tc.r)
			errNil := (err == nil)
			if errNil != tc.errNil {
				t.Errorf("Expected error == nil to be %v, but it was %v", tc.errNil, errNil)
			}
		})
	}
}

func TestSystemDims(t *testing.T) {
	tt := []struct {
		name string
		a    *mat.Dense
		b    *mat.Dense
		c    *mat.Dense
		q    *mat.Dense
		r    *mat.Dense
		pDim int
		mDim int
	}{
		{
			"1x1", mat.NewDense(1, 1, []float64{1}), mat.NewDense(1, 1, []float64{1}),
			mat.NewDense(1, 1, []float64{1}), mat.NewDense(1, 1, []float64{2}), mat.NewDense(1, 1, []float64{3}), 1, 1,
		},
		{
			"2x1", mat.NewDense(2, 2, []float64{1, 1, 0, 1}), mat.NewDense(2, 1, []float64{1, 1}),
			mat.NewDense(1, 2, []float64{1, 0}), mat.NewDense(1, 1, []float64{.5}), mat.NewDense(1, 1, []float64{.5}), 2, 1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			m, err := kalman.NewSystem(tc.a, tc.b, tc.c, tc.q, tc.r)
			if err != nil {
				t.Fatal("Failed to construct System")
			}
			pd, md := m.Dims()
			if pd != tc.pDim {
				t.Errorf("Expected process dim to be %v, but it was %v", tc.pDim, pd)
			}
			if md != tc.mDim {
				t.Errorf("Expected measurement dim to be %v, but it was %v", tc.mDim, md)
			}
		})
	}
}
