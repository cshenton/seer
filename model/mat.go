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
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// BlockDiag creates a block diagonal matrix with the provided matrices running
// along its diagonal.
func BlockDiag(mats ...*mat.Dense) (b *mat.Dense) {
	var (
		rowTotal  int
		colTotal  int
		colOffset int
		index     int
	)
	rows := make([]int, len(mats))
	cols := make([]int, len(mats))
	for i, m := range mats {
		r, c := m.Dims()
		rowTotal += r
		colTotal += c
		rows[i] = r
		cols[i] = c
	}

	data := make([]float64, rowTotal*colTotal)
	for i, m := range mats {
		for j := 0; j < rows[i]; j++ {
			// Prepend colOffset zeros
			for k := 0; k < colOffset; k++ {
				data[index] = 0
				index++
			}
			// Append the row's data
			for k := 0; k < cols[i]; k++ {
				data[index] = m.At(j, k)
				index++
			}
			// Postpend (colTotal - colOffset - cols[i]) zeros
			for k := 0; k < (colTotal - colOffset - cols[i]); k++ {
				data[index] = 0
				index++
			}
		}
		colOffset += cols[i]
	}
	b = mat.NewDense(rowTotal, colTotal, data)
	return b
}

// DenseValues extracts the row-major data from the provided dense matrix.
func DenseValues(d *mat.Dense) (data []float64) {
	nRows, _ := d.Dims()
	for i := 0; i < nRows; i++ {
		rowData := d.RawRowView(i)
		data = append(data, rowData...)
	}
	return data
}

// Diag returns the corresponding diagonal square matrix values given diag values.
func Diag(v []float64) (d []float64) {
	d = make([]float64, len(v)*len(v))
	for i := range v {
		d[i*(1+len(v))] = v[i]
	}
	return d
}

// Eye returns the dimension r identity matrix.
func Eye(r int) (m *mat.Dense, err error) {
	if r < 1 {
		err = fmt.Errorf("r must be 1 or greater, but was %v", r)
		return m, err
	}

	data := make([]float64, r*r)
	for i := range data {
		if i%(r+1) == 0 {
			data[i] = 1
		}
	}
	m = mat.NewDense(r, r, data)
	return m, nil
}
