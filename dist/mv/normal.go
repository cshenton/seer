package mv

import "errors"

// Normal is a multivariate normal distribution with full rank covariance.
type Normal struct {
	Location   []float64
	Covariance []float64
}

// NewNormal constructs a Normal with the given location, covariance (if possible).
func NewNormal(loc, cov []float64) (n *Normal, err error) {
	if len(loc) == 0 {
		err := errors.New("Location must have length 1 or greater")
		return nil, err
	}
	if len(cov) != len(loc)*len(loc) {
		err := errors.New("Covariance must have length equal to Location's squared length")
		return nil, err
	}
	n = &Normal{
		Location:   loc,
		Covariance: cov,
	}
	return n, nil
}

// Dim returns the dimensionality of the mv Normal distribution.
func (n *Normal) Dim() int {
	return len(n.Location)
}
