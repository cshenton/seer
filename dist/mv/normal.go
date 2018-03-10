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

package mv

import "errors"

// Normal is a multivariate normal distribution with full rank covariance.
type Normal struct {
	Location   []float64
	Covariance []float64
}

// NewNormal constructs a Normal with the given location, covariance (if possible).
func NewNormal(loc, cov []float64) (n *Normal, err error) {
	if len(loc) == 0 {
		err := errors.New("Location must have length 1 or greater")
		return nil, err
	}
	if len(cov) != len(loc)*len(loc) {
		err := errors.New("Covariance must have length equal to Location's squared length")
		return nil, err
	}
	n = &Normal{
		Location:   loc,
		Covariance: cov,
	}
	return n, nil
}

// Dim returns the dimensionality of the mv Normal distribution.
func (n *Normal) Dim() int {
	return len(n.Location)
}
