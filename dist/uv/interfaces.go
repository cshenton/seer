package uv

// Quantiler defines distributions that have computable quantiles. That is,
// which accept a probability on [0,1] and return a member in their support.
type Quantiler interface {
	Quantile(p float64) (q float64, err error)
}
