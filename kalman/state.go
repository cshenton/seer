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

package kalman

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// State describes the distribution of the internal state of a Linear system
// which has multivariate normal distribution.
type State struct {
	Loc *mat.Dense
	Cov *mat.Dense
}

// NewState returns a State if the provided matrices have compatible dimensions
// and otherwise returns an error.
func NewState(loc, cov *mat.Dense) (s *State, err error) {
	lRow, lCol := loc.Dims()
	cRow, cCol := cov.Dims()
	if lCol != 1 {
		err = fmt.Errorf("loc must be a column matrix, but had %v columns", lCol)
		return s, err
	}
	if cRow != cCol {
		err = fmt.Errorf("cov must be a square matrix, but was %v by %v", cRow, cCol)
		return s, err
	}
	if lRow != cRow {
		err = fmt.Errorf("loc and cov must have same number of rows, but had %v and %v", lRow, cRow)
		return s, err
	}
	s = &State{loc, cov}
	return s, nil
}

// Dim returns the integer dimension of the state.
func (s *State) Dim() (d int) {
	d, _ = s.Cov.Dims()
	return d
}
