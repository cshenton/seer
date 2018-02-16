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

	det := m.Deterministic.Forecast(period, n)
	stoch := m.Stochastic.Forecast(m.RCE.Noise(), m.RCE.Walk(), n)

	for i := range f {
		d := &uv.Normal{
			Location: det[i].Location + stoch[i].Location,
			Scale:    math.Sqrt(math.Pow(det[i].Scale, 2) + math.Pow(stoch[i].Scale, 2)),
		}
		f[i] = d
	}
	return f
}
