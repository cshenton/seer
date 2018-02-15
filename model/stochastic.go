package model

import (
	"github.com/chulabs/seer/dist/mv"
	"github.com/chulabs/seer/kalman"
	"gonum.org/v1/gonum/mat"
)

// Stochastic is the type against which we apply stochastic model updates.
type Stochastic struct {
	*mv.Normal
}

// NewStochastic returns a stochastic with a proper state prior.
func NewStochastic() (s *Stochastic) {
	loc := []float64{0}
	cov := []float64{1e12}
	n, _ := mv.NewNormal(loc, cov)
	s = &Stochastic{n}
	return s
}

// System generates process and observation matrices for this linear system.
func (s *Stochastic) System(noise, walk float64) (k *kalman.System) {
	a := mat.NewDense(1, 1, []float64{1})
	b := mat.NewDense(1, 1, []float64{1})
	c := mat.NewDense(1, 1, []float64{1})
	q := mat.NewDense(1, 1, []float64{walk})
	r := mat.NewDense(1, 1, []float64{noise})

	k, _ = kalman.NewSystem(a, b, c, q, r)
	return k
}

// Update performs a filter step against the stochastic state.
func (s *Stochastic) Update(noise, walk, val float64) {
	// p, pc, o, oc := s.Filters(noise, walk)
}
