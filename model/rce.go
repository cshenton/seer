package model

import (
	"math"

	"github.com/cshenton/seer/dist/uv"
)

// History contains the last two observations for this stream, it is required
// to do covariance estimation in real time.
type History []float64

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

// NewRCE constructs an RCE with appropriate priors for theta and zeta.
func NewRCE() (r *RCE) {
	h := History([]float64{0, 0})
	t, _ := uv.NewInverseGamma(2, 180)
	z, _ := uv.NewInverseGamma(2, 100)

	r = &RCE{
		Theta:   t,
		Zeta:    z,
		History: h,
	}
	return r
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
	r.Zeta.Shape += 0.5
	r.Theta.Shape += 0.5
	r.Zeta.Scale += math.Pow(v-r.History[0], 2) / 2
	r.Theta.Scale += math.Pow(v-r.History[1], 2) / 2
	r.History.Update(v)
}
