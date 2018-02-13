package model

// methods for filtering across rce, det, stoch

// get covariances from RCE
// construct det filters using period and sum noise, walk covariance as det noise
// forward pass deterministic filter, get residual, state
// update rce on residual
// forward pass stochastic filter, state
// increment time
// apply new time
// apply new states
// apply new RCE

// Filter filters
func Filter(v, p float64, d *Deterministic, s *Stochastic, r *RCE) {
	r.Update(v)
	noise := r.Noise()
	walk := r.Walk()
	resid := d.Update(noise, walk, p, v)
	s.Update(noise, walk, resid)
}

// Forecast forecasts
func Forecast(n int, p float64, d *Deterministic, s *Stochastic, r *RCE) {
	// noise := r.Noise()
	// walk := r.Walk()

	// Forecasts without obs noise
	// dets := d.Forecast(noise, walk, p, n)
	// Forecasts with obs noise
	// stochs := s.Forecast(noise, walk, n)

	// Add the dets, stochs together
	// Return a slice of
}
