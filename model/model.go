package model

import (
	"time"
)

// Model stores dynamic state about a stream.
type Model struct {
	Time          time.Time
	Deterministic *Deterministic
	Stochastic    *Stochastic
	RCE           *RCE
}

// New initialises a model given a stream period.
func New(period float64) (m *Model) {
	m = &Model{
		Deterministic: NewDeterministic(period),
		Stochastic:    NewStochastic(),
		RCE:           NewRCE(),
	}
	return m
}

// Update iterates the Model in response to an observed event.
func (m *Model) Update(v, period float64) {
	// get covariances from RCE
	// construct filters using period (and external package)
	// forward pass deterministic filter, get residual, state
	// update rce on residual
	// forward pass stochastic filter, state
	// increment time
	// apply new time
	// apply new states
	// apply new RCE
}

// func Filter(v, p float64, d *Deterministic, s *Stochastic, r *RCE) {
// 	// Update d, get resid
// 	// update r on resid
// 	// update s with updated r
// 	r.Update(v)
// 	noise := r.Noise()
// 	walk := r.Walk()
// 	resid := d.Update(noise, walk, p, v)
// 	s.Update(noise, walk, resid)
// }

// Forecast passes forward through the filter and returns
