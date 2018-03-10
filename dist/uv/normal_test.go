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

package uv_test

import (
	"math"
	"testing"

	"github.com/cshenton/seer/dist/uv"
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
	q, err := n.Quantile(0.5)
	if err != nil {
		t.Error("unexpected error in Quantile,", err)
	}
	if math.Abs(q-loc) > 1e-8 {
		t.Errorf("expected median quantile %v, but got %v", loc, q)
	}
}

func TestNormalQuantileErrs(t *testing.T) {
	loc := 0.0
	scale := 1.0

	n, err := uv.NewNormal(loc, scale)
	if err != nil {
		t.Error("unexpected error in NewNormal,", err)
	}
	_, err = n.Quantile(2)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}
