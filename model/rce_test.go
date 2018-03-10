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

func TestHistoryUpdate(t *testing.T) {
	h := model.History([]float64{1, 2})
	h.Update(3)

	if h[0] != 3 {
		t.Errorf("expected entry 0 to be %v, but it was %v", 3, h[0])
	}
	if h[1] != 1 {
		t.Errorf("expected entry 1 to be %v, but it was %v", 1, h[1])
	}
}

func TestNewRCE(t *testing.T) {
	r := model.NewRCE()

	if r.History[0] != 0 {
		t.Errorf("expected history 0 of %v, but it was %v", 0, r.History[0])
	}
	if r.History[1] != 0 {
		t.Errorf("expected history 1 of %v, but it was %v", 0, r.History[1])
	}
	if r.Theta.Scale != 180 {
		t.Errorf("expected theta scale of %v, but it was %v", 180, r.Theta.Scale)
	}
	if r.Zeta.Scale != 100 {
		t.Errorf("expected zeta scale of %v, but it was %v", 100, r.Zeta.Scale)
	}
}

func TestRCEWalk(t *testing.T) {
	r := model.NewRCE()

	if r.Walk() != 10.0 {
		t.Errorf("expected walk covariance of %v, but it was %v", 10.0, r.Walk())
	}
}

func TestRCENoise(t *testing.T) {
	r := model.NewRCE()

	if r.Noise() != 80.0 {
		t.Errorf("expected Noise covariance of %v, but it was %v", 80.0, r.Noise())
	}
}

func TestRCEUpdate(t *testing.T) {
	r := model.NewRCE()
	r.Update(1.0)

	if r.Theta.Shape != 2.5 {
		t.Errorf("expected theta shape %v, got %v", 2.5, r.Theta.Shape)
	}
	if r.Theta.Scale != 180.5 {
		t.Errorf("expected theta scale %v, got %v", 180.5, r.Theta.Scale)
	}
	if r.Zeta.Shape != 2.5 {
		t.Errorf("expected zeta shape %v, got %v", 2.5, r.Zeta.Shape)
	}
	if r.Zeta.Scale != 100.5 {
		t.Errorf("expected theta scale %v, got %v", 100.5, r.Zeta.Scale)
	}
	if r.History[0] != 1 {
		t.Errorf("expected history 0 of %v, got %v", 1, r.History[0])
	}
	if r.History[1] != 0 {
		t.Errorf("expected history 1 of %v, got %v", 0, r.History[1])
	}
}
