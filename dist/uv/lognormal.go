package uv

import "errors"

// LogNormal is the log normal distribution.
type LogNormal struct {
	Location float64
	Scale    float64
}

// NewLogNormal creates a new lognormal distribution, ensuring first that
// the provided scale is strictly positive.
func NewLogNormal(location, scale float64) (ln *LogNormal, err error) {
	if scale <= 0 {
		err := errors.New("scale must be strictly greater than zero")
		return nil, err
	}
	ln = &LogNormal{
		Location: location,
		Scale:    scale,
	}
	return ln, nil
}
