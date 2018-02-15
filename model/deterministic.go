package model

import (
	"math"

	"github.com/chulabs/seer/dist/mv"
	"github.com/chulabs/seer/kalman"
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

// System generates process and observation matrices for this linear system.
func (d *Deterministic) System(period, noise, walk float64) (k *kalman.System) {
	h := Harmonics(period, maxHarmonic)
	dim := 2*len(h) + 2

	a := mat.NewDense(1, 1, []float64{noise + walk})

	b := mat.NewDense(1, 1, []float64{noise + walk})

	c := mat.NewDense(1, 1, []float64{noise + walk})

	qVals := make([]float64, dim)
	qVals[0] = levelVar
	qVals[1] = trendVar
	for i := 2; i < len(qVals); i++ {
		qVals[i] = harmonicVar
	}
	q := mat.NewDense(dim, dim, Diag([]float64{}))

	r := mat.NewDense(1, 1, []float64{noise + walk})

	k, _ = kalman.NewSystem(a, b, c, q, r)
	return k
	// A: block diag of 1101 and the harmonics
	// B: (eye dim)
	// C: 101010 dim
	// Q: diag (level trend harmonic...)
	// R: 1x1 filter noise
	// make system
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
