package dist

import "math"

// InverseGamma is the inverse gamma distribution.
type InverseGamma struct {
	Shape float64
	Scale float64
}

// UpdateNormal applies a conjugate prior update where ig is the scale
// distribution, v is the observed value, and m is the normal mean.
func (i *InverseGamma) UpdateNormal(v, m float64) {
	i.Shape += 0.5
	i.Scale += math.Pow(v-m, 2) / 2
}

// Mean returns the first moment of the distribution, for shape > 1.
func (i *InverseGamma) Mean() float64 {
	return i.Scale / (i.Shape - 1)
}

// Variance returns the second central moment of the distribution, for shape > 2.
func (i *InverseGamma) Variance() float64 {
	return math.Pow(i.Scale, 2) / (math.Pow(i.Shape-1, 2) * (i.Shape - 2))
}
