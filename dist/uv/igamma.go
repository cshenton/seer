package uv

import (
	"errors"
	"math"
)

// InverseGamma is the inverse gamma distribution.
type InverseGamma struct {
	Shape float64
	Scale float64
}

// NewInverseGamma constructs an InverseGamma, validating the provided parameters.
func NewInverseGamma(shape, scale float64) (i *InverseGamma, err error) {
	if shape <= 0.0 || scale <= 0.0 {
		err := errors.New("shape and scale must both be strictly positive")
		return nil, err
	}
	i = &InverseGamma{
		Shape: shape,
		Scale: scale,
	}
	return i, nil
}

// Mean returns the first moment of the distribution, for shape > 1.
func (i *InverseGamma) Mean() float64 {
	return i.Scale / (i.Shape - 1)
}

// Variance returns the second central moment of the distribution, for shape > 2.
func (i *InverseGamma) Variance() float64 {
	return math.Pow(i.Scale, 2) / (math.Pow(i.Shape-1, 2) * (i.Shape - 2))
}
