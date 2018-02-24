package uv

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mathext"
)

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

// Quantile is the inverse function of the log normal CDF.
func (ln *LogNormal) Quantile(p float64) float64 {
	return math.Exp(ln.Location + ln.Scale*mathext.NormalQuantile(p))
}
