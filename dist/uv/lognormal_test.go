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

func TestLogNormalQuantile(t *testing.T) {
	loc := 0.0
	scale := 1.0

	ln, err := uv.NewLogNormal(loc, scale)
	if err != nil {
		t.Error("unexpected error in NewLogNormal,", err)
	}
	q, err := ln.Quantile(0.5)
	if err != nil {
		t.Error("unexpected error in Quantile,", err)
	}
	if math.Abs(q-math.Exp(loc)) > 1e-8 {
		t.Errorf("expected median quantile %v, but got %v", math.Exp(loc), q)
	}
}

func TestLogNormalQuantileErrs(t *testing.T) {
	loc := 0.0
	scale := 1.0

	ln, err := uv.NewLogNormal(loc, scale)
	if err != nil {
		t.Error("unexpected error in NewLogNormal,", err)
	}
	_, err = ln.Quantile(2)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}
