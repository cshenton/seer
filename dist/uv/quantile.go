package uv

// Quantiler defines distributions that have computable quantiles. That is,
// which accept a probability on [0,1] and return a member in their support.
type Quantiler interface {
	Quantile(p float64) (q float64, err error)
}

// ConfidenceInterval constructs a confidence interval with confidence level p,
// where higher confidence in [0,1] means a wider interval.
func ConfidenceInterval(q Quantiler, p float64) (l, u float64, err error) {
	l, err = q.Quantile(0.5 - p/2)
	if err != nil {
		return 0, 0, err
	}
	u, err = q.Quantile(0.5 + p/2)
	if err != nil {
		return 0, 0, err
	}
	return l, u, nil
}
