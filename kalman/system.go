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

// System contains a set of transition and covariance matrices which fully
// describe a Linear Time Invariant (LTI) dynamical system.
type System struct {
	// Process matrix
	A *mat.Dense
	// Process noise transform matrix
	B *mat.Dense
	// Measurement matrix
	C *mat.Dense
	// Process Covariance matrix
	Q *mat.Dense
	// Measurement Covariance matrix
	R *mat.Dense
}

// NewSystem returns a System if the provided matrices have compatible dimensions
// and otherwise returns an error.
func NewSystem(a, b, c, q, r *mat.Dense) (m *System, err error) {
	aRow, aCol := a.Dims()
	bRow, bCol := b.Dims()
	cRow, cCol := c.Dims()
	qRow, qCol := q.Dims()
	rRow, rCol := r.Dims()
	if qRow != qCol {
		err = fmt.Errorf("Q must be square, but was %v, by %v", qRow, qCol)
		return m, err
	}
	if rRow != rCol {
		err = fmt.Errorf("R must be square, but was %v, by %v", rRow, rCol)
		return m, err
	}
	if aRow != aCol {
		err = fmt.Errorf("A must be square, but was %v, by %v", aRow, aCol)
		return m, err
	}
	if aRow != bRow {
		err = fmt.Errorf("A rows and B rows must match but were %v and %v", aRow, bRow)
		return m, err
	}
	if bCol != qCol {
		err = fmt.Errorf("B cols and Q rows must match but were %v and %v", bCol, qCol)
		return m, err
	}
	if cRow != rRow {
		err = fmt.Errorf("C rows and R rows must match, but were %v and %v", cRow, rRow)
		return m, err
	}
	if cCol != aRow {
		err = fmt.Errorf("C cols must match A rows, but were %v and %v", cCol, aRow)
		return m, err
	}
	m = &System{a, b, c, q, r}
	return m, nil
}

// Dims returns the dimensions of the process and measurement equations of the
// given linear System.
func (m *System) Dims() (pDim, mDim int) {
	pDim, _ = m.A.Dims()
	mDim, _ = m.C.Dims()
	return pDim, mDim
}
