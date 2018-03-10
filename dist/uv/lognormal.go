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

package uv

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mathext"
)

// LogNormal is the log normal distribution.
type LogNormal struct {
	Location float64
	Scale    float64
}

// NewLogNormal creates a new lognormal distribution, ensuring first that
// the provided scale is strictly positive.
func NewLogNormal(location, scale float64) (ln *LogNormal, err error) {
	if scale <= 0 {
		err := errors.New("scale must be strictly greater than zero")
		return nil, err
	}
	ln = &LogNormal{
		Location: location,
		Scale:    scale,
	}
	return ln, nil
}

// Quantile is the inverse function of the log normal CDF.
func (ln *LogNormal) Quantile(p float64) (q float64, err error) {
	if p < 0 || p > 1 {
		err := errors.New("probabilities must be between 0 and 1")
		return q, err
	}
	q = math.Exp(ln.Location + ln.Scale*mathext.NormalQuantile(p))
	return q, nil
}
