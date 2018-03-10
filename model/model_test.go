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

func TestNew(t *testing.T) {
	m := model.New(604800)

	if len(m.Deterministic.Location) != 22 {
		t.Errorf("expected deterministic location dim of %v, but got %v", 22, len(m.Deterministic.Location))
	}
	if len(m.Stochastic.Location) != 1 {
		t.Errorf("expected stochastic location dim of %v, but got %v", 1, len(m.Stochastic.Location))
	}
	if m.RCE.Theta.Scale != 180 {
		t.Errorf("expected initial theta scale of %v, but got %v", 180, m.RCE.Theta.Scale)
	}
}

func TestModelUpdate(t *testing.T) {
	m := model.New(604800)

	m.Update(604800, 1.0)

	if m.Deterministic.Location[0] == 0 {
		t.Error("deterministic not updated")
	}

	if m.Stochastic.Location[0] == 0 {
		t.Error("stochastic not updated")
	}

	if m.RCE.History[0] == 0 {
		t.Error("RCE not updated")
	}
}

func TestModelForecast(t *testing.T) {
	period := 604800.0
	n := 150

	m := model.New(period)
	m.Update(period, 1)
	f := m.Forecast(period, n)

	if len(f) != n {
		t.Errorf("expected length %v, but it was %v", n, len(f))
	}
}
