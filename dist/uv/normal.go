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

// Normal distribution defined by its mean and standard deviation.
type Normal struct {
	Location float64
	Scale    float64
}

// NewNormal checks the input parameters and returns a Normal constructed
// using them, if they are valid.
func NewNormal(location, scale float64) (n *Normal, err error) {
	if scale <= 0 {
		err := errors.New("scale must be strictly greater than zero")
		return nil, err
	}
	n = &Normal{
		Location: location,
		Scale:    scale,
	}
	return n, nil
}

// Mean returns the first moment of the distribution.
func (n *Normal) Mean() float64 {
	return n.Location
}

// Variance returns the second central moment of the distribution.
func (n *Normal) Variance() float64 {
	return math.Pow(n.Scale, 2)
}

// Quantile is the inverse function of the CDF.
func (n *Normal) Quantile(p float64) (q float64, err error) {
	if p < 0 || p > 1 {
		err := errors.New("probabilities must be between 0 and 1")
		return q, err
	}
	q = n.Location + n.Scale*mathext.NormalQuantile(p)
	return q, nil
}
