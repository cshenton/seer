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

package model_test

import (
	"testing"

	"github.com/cshenton/seer/model"
)

func TestNewStochastic(t *testing.T) {
	s := model.NewStochastic()

	if s.Location[0] != 0.0 {
		t.Error("Expected location 0.0 but got", s.Location[0])
	}
	if s.Covariance[0] != 1e12 {
		t.Error("Expected covariance 1e12, but got", s.Covariance[0])
	}
}

func TestStochasticState(t *testing.T) {
	s := model.NewStochastic()

	k := s.State()

	lx, ly := k.Loc.Dims()
	if lx != 1 {
		t.Fatalf("expected location to have %v rows, but it had %v", 1, lx)
	}
	if ly != 1 {
		t.Fatalf("expected location to have %v columns, but it had %v", 1, ly)
	}

	cx, cy := k.Cov.Dims()
	if cx != 1 {
		t.Fatalf("expected covariance to have %v rows, but it had %v", 1, cx)
	}
	if cy != 1 {
		t.Fatalf("expected covariance to have %v columns, but it had %v", 1, cy)
	}
}

func TestStochasticSystem(t *testing.T) {
	s := model.NewStochastic()
	k := s.System(100, 10)

	if k.Q.At(0, 0) != 10 {
		t.Errorf("expected process covariance of %v, but it was %v", 10, k.Q.At(0, 0))
	}
	if k.R.At(0, 0) != 100 {
		t.Errorf("expected observation covariance of %v, but it was %v", 100, k.R.At(0, 0))
	}
}

func TestStochasticUpdate(t *testing.T) {
	s := model.NewStochastic()

	err := s.Update(100, 10, 1)

	if err != nil {
		t.Fatal("unexpected error during Update:", err)
	}

	if s.Location[0] == 0 {
		t.Error("location appears not to be updated")
	}
}

func TestStochasticForecast(t *testing.T) {
	noise := 100.0
	walk := 10.0
	n := 100

	s := model.NewStochastic()
	err := s.Update(100, 10, 1)
	if err != nil {
		t.Fatal("unexpected error in Update:", err)
	}

	f := s.Forecast(noise, walk, n)

	if len(f) != n {
		t.Errorf("expected length %v, but it was %v", n, len(f))
	}
}
