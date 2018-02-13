package model

import (
	"github.com/chulabs/seer/dist/mv"
	"gonum.org/v1/gonum/mat"
)

// Stochastic is the type against which we apply stochastic model updates.
type Stochastic struct {
	*mv.Normal
}

// Filters generates process and observation matrices for this linear system.
func (s *Stochastic) Filters(noise, walk float64) (p, pc, o, oc *mat.Dense) {
	p = mat.NewDense(1, 1, []float64{1})
	pc = mat.NewDense(1, 1, []float64{walk})
	o = mat.NewDense(1, 1, []float64{1})
	oc = mat.NewDense(1, 1, []float64{noise})
	return p, pc, o, oc
}

// Update performs a filter step against the stochastic state.
func (s *Stochastic) Update(noise, walk, val float64) {
	// p, pc, o, oc := s.Filters(noise, walk)
}
