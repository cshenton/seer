package uv

import (
	"math"

	"gonum.org/v1/gonum/mathext"
)

// Normal distribution defined by its mean and standard deviation.
type Normal struct {
	Location float64
	Scale    float64
}

// Mean returns the first moment of the distribution.
func (n *Normal) Mean() float64 {
	return n.Location
}

// Variance returns the second central moment of the distribution.
func (n *Normal) Variance() float64 {
	return math.Pow(n.Scale, 2)
}

// Quantile is the inverse function of the CDF.
func (n *Normal) Quantile(p float64) float64 {
	return n.Location + n.Scale*mathext.NormalQuantile(p)
}
