package model

import (
	"math"

	"github.com/chulabs/seer/dist/uv"
)

// History contains the last two observations for this stream, it is required
// to do covariance estimation in real time.
type History [2]float64

// Update shifts the history one place down and adds a new value.
func (h History) Update(v float64) {
	h[1] = h[0]
	h[0] = v
}

// RCE is a recursive covariance estimator for a local level kalman filter.
type RCE struct {
	Theta   *uv.InverseGamma
	Zeta    *uv.InverseGamma
	History History
}

// Walk returns the current walk covariance.
func (r *RCE) Walk() float64 {
	return math.Abs(r.Zeta.Mean() - 0.5*r.Theta.Mean())
}

// Noise returns the current noise covariance.
func (r *RCE) Noise() float64 {
	return math.Abs(r.Zeta.Mean() - 2*r.Walk())
}

// Update updates the covariance estimator.
func (r *RCE) Update(v float64) {
	r.History.Update(v)
	r.Zeta.Update(v-r.History[0], 0)
	r.Theta.Update(v-r.History[1], 0)
}
