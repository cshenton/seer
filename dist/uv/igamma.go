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
)

// InverseGamma is the inverse gamma distribution.
type InverseGamma struct {
	Shape float64
	Scale float64
}

// NewInverseGamma constructs an InverseGamma, validating the provided parameters.
func NewInverseGamma(shape, scale float64) (i *InverseGamma, err error) {
	if shape <= 0.0 || scale <= 0.0 {
		err := errors.New("shape and scale must both be strictly positive")
		return nil, err
	}
	i = &InverseGamma{
		Shape: shape,
		Scale: scale,
	}
	return i, nil
}

// Mean returns the first moment of the distribution, for shape > 1.
func (i *InverseGamma) Mean() float64 {
	return i.Scale / (i.Shape - 1)
}

// Variance returns the second central moment of the distribution, for shape > 2.
func (i *InverseGamma) Variance() float64 {
	return math.Pow(i.Scale, 2) / (math.Pow(i.Shape-1, 2) * (i.Shape - 2))
}
