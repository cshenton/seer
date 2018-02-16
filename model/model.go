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
func (m *Model) Update(period, val float64) {
	resid, _ := m.Deterministic.Update(m.RCE.Noise(), m.RCE.Walk(), period, val)
	m.RCE.Update(resid)
	m.Stochastic.Update(m.RCE.Noise(), m.RCE.Walk(), resid)
}
