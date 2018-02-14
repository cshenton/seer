package model

import (
	"math"

	"github.com/chulabs/seer/dist/mv"
	"gonum.org/v1/gonum/mat"
)

const (
	maxHarmonic = 31577600
	harmonicVar = 1e4
	levelVar    = 1e15
	trendVar    = 1e5
)

// Harmonics provides a consistent method for generating fourier harmonics for
// a stream, it does so by splitting harmonics into powers of 10.
func Harmonics(min, max float64) []float64 {
	ratio := max / min
	mag := int(math.Log10(ratio))
	harmonics := make([]float64, mag*9+1)

	for i := 0; i < mag; i++ {
		for j := 1; j < 10; j++ {
			harmonics[9*i+j-1] = (max / math.Pow(10, float64(i))) / float64(j)
		}
	}
	harmonics[9*mag] = max / math.Pow(10, float64(mag))
	return harmonics
}

// Diag returns the corresponding diagonal square matrix values given diag values.
func Diag(v []float64) (d []float64) {
	d = make([]float64, len(v)*len(v))
	for i := range v {
		d[i*(1+len(v))] = v[i]
	}
	return d
}

// Deterministic is the type against which we apply deterministic model updates.
type Deterministic struct {
	*mv.Normal
}

// NewDeterministic creates and returns a Deterministic with a proper state prior.
func NewDeterministic(period float64) (d *Deterministic) {
	dim := 2*len(Harmonics(period, maxHarmonic)) + 2
	loc := make([]float64, dim)
	v := make([]float64, dim)
	v[0] = levelVar
	v[1] = trendVar
	for i := 2; i < dim; i++ {
		v[i] = harmonicVar
	}
	cov := Diag(v)
	n, _ := mv.NewNormal(loc, cov)
	d = &Deterministic{n}
	return d
}

// Filters generates process and observation matrices for this linear system.
func (d *Deterministic) Filters(period float64) (p, pc, o, oc *mat.Dense) {
	// make deterministic system from period, noise param
	// return
	return
}

// Update performs a filter step against the deterministic state.
func (d *Deterministic) Update(noise, walk, period, val float64) float64 {
	// Get System from period, noise param
	// Filter and get residual
	// update state
	// return residual

	// p, pc, o, oc := s.Filters(noise, walk)
	return 1.0
}

// method for generating process, proc cov, obs, obs cov matrices

// method for updating deterministic, returning residual
