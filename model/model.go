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

package model

import (
	"math"

	"github.com/cshenton/seer/dist/uv"
)

// Model stores dynamic state about a stream.
type Model struct {
	Deterministic *Deterministic
	Stochastic    *Stochastic
	RCE           *RCE
}

// New initialises a model given a stream period.
func New(period float64) (m *Model) {
	m = &Model{
		Deterministic: NewDeterministic(period),
		Stochastic:    NewStochastic(),
		RCE:           NewRCE(),
	}
	return m
}

// Update iterates the Model in response to an observed event.
func (m *Model) Update(period, val float64) {
	resid, _ := m.Deterministic.Update(m.RCE.Noise(), m.RCE.Walk(), period, val)
	m.RCE.Update(resid)
	m.Stochastic.Update(m.RCE.Noise(), m.RCE.Walk(), resid)
}

// Forecast returns a slice of Normally distributed predictions.
func (m *Model) Forecast(period float64, n int) (f []*uv.Normal) {
	f = make([]*uv.Normal, n)

	d := m.Deterministic.Forecast(period, n)
	s := m.Stochastic.Forecast(m.RCE.Noise(), m.RCE.Walk(), n)

	for i := range f {
		dist := &uv.Normal{
			Location: d[i].Location + s[i].Location,
			Scale:    math.Sqrt(math.Pow(d[i].Scale, 2) + math.Pow(s[i].Scale, 2)),
		}
		f[i] = dist
	}
	return f
}
