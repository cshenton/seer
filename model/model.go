package model

import (
	"math"

	"github.com/chulabs/seer/dist/uv"
)

// Model stores dynamic state about a stream.
type Model struct {
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

// Forecast returns a slice of Normally distributed predictions.
func (m *Model) Forecast(period float64, n int) (f []*uv.Normal) {
	f = make([]*uv.Normal, n)

	d := m.Deterministic.Forecast(period, n)
	s := m.Stochastic.Forecast(m.RCE.Noise(), m.RCE.Walk(), n)

	for i := range f {
		dist := &uv.Normal{
			Location: d[i].Location + s[i].Location,
			Scale:    math.Sqrt(math.Pow(d[i].Scale, 2) + math.Pow(s[i].Scale, 2)),
		}
		f[i] = dist
	}
	return f
}
