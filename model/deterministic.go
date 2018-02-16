package model

import (
	"math"

	"github.com/chulabs/seer/dist/mv"
	"github.com/chulabs/seer/dist/uv"
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

// State returns the kalman filter State.
func (d *Deterministic) State() (k *kalman.State) {
	l := mat.NewDense(d.Dim(), 1, d.Location)
	c := mat.NewDense(d.Dim(), d.Dim(), d.Covariance)

	k, _ = kalman.NewState(l, c)
	return k
}

// System generates process and observation matrices for this linear system.
func (d *Deterministic) System(noise, walk, period float64) (k *kalman.System) {
	h := Harmonics(period, maxHarmonic)
	dim := 2*len(h) + 2

	aMats := make([]*mat.Dense, len(h)+1)
	aMats[0] = mat.NewDense(2, 2, []float64{1, 1, 0, 1})
	for i := 1; i < len(aMats); i++ {
		angle := 2 * math.Pi / h[i-1]
		data := []float64{math.Cos(angle), math.Sin(angle), -math.Sin(angle), math.Cos(angle)}
		aMats[i] = mat.NewDense(2, 2, data)
	}
	a := BlockDiag(aMats...)

	b, _ := Eye(dim)

	cVals := make([]float64, dim)
	for i := range cVals {
		if i%2 == 0 {
			cVals[i] = 1
		}
	}
	c := mat.NewDense(1, dim, cVals)

	qVals := make([]float64, dim)
	qVals[0] = levelVar
	qVals[1] = trendVar
	for i := 2; i < len(qVals); i++ {
		qVals[i] = harmonicVar
	}
	q := mat.NewDense(dim, dim, Diag(qVals))

	r := mat.NewDense(1, 1, []float64{noise + walk})

	k, _ = kalman.NewSystem(a, b, c, q, r)
	return k
}

// Update performs a filter step against the deterministic state.
func (d *Deterministic) Update(noise, walk, period, val float64) (resid float64, err error) {
	st := d.State()
	sy := d.System(noise, walk, period)

	statePred, _ := kalman.Predict(st, sy)
	newState, resid, _ := kalman.Update(statePred, sy, val)

	d.Location = DenseValues(newState.Loc)
	d.Covariance = DenseValues(newState.Cov)
	return resid, nil
}

// Forecast returns a forecasted slice of normal RVs for this deterministic component.
func (d *Deterministic) Forecast(period float64, n int) (f []*uv.Normal) {
	f = make([]*uv.Normal, n)

	st := d.State()
	sy := d.System(0, 0, period)

	for i := 0; i < n; i++ {
		st, _ = kalman.Predict(st, sy)
		ob, _ := kalman.Observe(st, sy)
		f[i] = &uv.Normal{
			Location: DenseValues(ob.Loc)[0],
			Scale:    math.Sqrt(DenseValues(ob.Cov)[0]),
		}
	}
	return f
}
