package kalman

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Kalman defines the set of standard Kalman filter update equations on a
// linear system. It is an empty struct, rather than a package, so that it
// has parity with other concrete implementations of filter.Filter that, unlike
// kalman.Filter, are stateful (for example kalman.FilterRCE).
type Kalman struct {
	State  *State
	System *System
}

// New creates a filter from the given state, system, and validates that
// their dimensions match.
func New(st *State, sys *System) (k *Kalman, err error) {
	stDim := st.Dim()
	sysDim, _ := sys.Dims()
	if stDim != sysDim {
		err = fmt.Errorf("State dim must match System process dim, but were %v, and %v", stDim, sysDim)
		return k, err
	}
	k = &Kalman{st, sys}
	return k, nil
}

// Filter updates the internal state by doing the predict and update steps of
// the kalman filter. It returns the post-fit residual, and an error if any
// dimensions inside the filter do not match.
func (k *Kalman) Filter(v float64) (res float64, err error) {
	statePred, err := Predict(k.State, k.System)
	if err != nil {
		return res, err
	}
	stateNew, res, err := Update(statePred, k.System, v)
	if err != nil {
		return res, err
	}
	k.State = stateNew
	return res, nil
}

// Predict predicts the next state distribution given the previous state and
// linear system equation `x_next = A * x_prev + B * w_prev `.
func Predict(p *State, m *System) (n *State, err error) {
	var loc mat.Dense
	var cov mat.Dense
	var addcov mat.Dense

	pDim := p.Dim()
	mDim, _ := m.Dims()

	if pDim != mDim {
		err = fmt.Errorf("Prev state dim must match process dim, but were %v, and %v", pDim, mDim)
		return n, err
	}

	loc.Mul(m.A, p.Loc)
	cov.Product(m.A, p.Cov, m.A.T())
	addcov.Product(m.B, m.Q, m.B.T())
	cov.Add(&cov, &addcov)
	n, err = NewState(&loc, &cov)
	return n, err
}

// Observe determines the observation distribution for the current time step
// using the linear measurement equation: `y_next = C * x_next + v_next`. It
// returns an error if the state and system process dims do not match.
func Observe(s *State, m *System) (o *State, err error) {
	var loc mat.Dense
	var cov mat.Dense

	mDim, _ := m.Dims()
	sDim := s.Dim()
	if sDim != mDim {
		err = fmt.Errorf("Prev state dim must match process dim, but were %v, and %v", sDim, mDim)
		return o, err
	}
	loc.Mul(m.C, s.Loc)
	cov.Product(m.C, s.Cov, m.C.T())
	cov.Add(&cov, m.R)
	o, err = NewState(&loc, &cov)
	return o, err
}

// StateObserve foo
func StateObserve(s *State, m *System) (o *State, err error) {
	var loc mat.Dense
	var cov mat.Dense

	mDim, _ := m.Dims()
	sDim := s.Dim()
	if sDim != mDim {
		err = fmt.Errorf("Prev state dim must match process dim, but were %v, and %v", sDim, mDim)
		return o, err
	}
	loc.Mul(m.C, s.Loc)
	cov.Product(m.C, s.Cov, m.C.T())
	o, err = NewState(&loc, &cov)
	return o, err
}

// Update implements the Kalman Filter update, and returns the a posteriori
// internal state distribution given the state prior and an observation.
// Also returns the post-fit residual.
func Update(p *State, m *System, v float64) (n *State, res float64, err error) {
	var (
		ir    mat.Dense
		ic    mat.Dense
		ici   mat.Dense
		gain  mat.Dense
		loc   mat.Dense
		cov   mat.Dense
		resid mat.Dense
	)
	pDim := p.Dim()
	mDim, _ := m.Dims()

	if pDim != mDim {
		err = fmt.Errorf("State dim must match process dim, but were %v, and %v", pDim, mDim)
		return n, res, err
	}

	ir.Mul(m.C, p.Loc)
	ir.Scale(-1.0, &ir)
	ir.Apply(func(i, j int, val float64) float64 { return val + v }, &ir)
	ic.Product(m.C, p.Cov, m.C.T())
	ic.Add(&ic, m.R)
	ici.Inverse(&ic)
	gain.Product(p.Cov, m.C.T(), &ici)
	loc.Mul(&gain, &ir)
	loc.Add(p.Loc, &loc)
	cov.Product(&gain, &ic, gain.T())
	cov.Scale(-1.0, &cov)
	cov.Add(p.Cov, &cov)
	resid.Mul(m.C, &loc)
	res = v - resid.At(0, 0)
	n, _ = NewState(&loc, &cov)
	return n, res, err
}
