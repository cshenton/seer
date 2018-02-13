package kalman

import (
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
