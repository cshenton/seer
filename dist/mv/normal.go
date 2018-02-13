package mv

// Normal is a multivariate normal distribution with full rank covariance.
type Normal struct {
	Dim        int
	Location   []float64
	Covariance []float64
}
