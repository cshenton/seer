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
