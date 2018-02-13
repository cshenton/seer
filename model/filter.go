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
