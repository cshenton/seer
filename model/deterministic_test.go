package model_test

import (
	"testing"

	"github.com/cshenton/seer/model"
)

// yr, seconds in an average year (365.25 days)
const yr = float64(31577600)

func TestHarmonics(t *testing.T) {
	tt := []struct {
		name   string
		min    float64
		max    float64
		result []float64
	}{
		{
			"1 and 10", 1, 10, []float64{
				10.0, 10.0 / 2, 10.0 / 3, 10.0 / 4, 10.0 / 5, 10.0 / 6, 10.0 / 7, 10.0 / 8,
				10.0 / 9, 10.0 / 10.0,
			},
		},
		{
			"1 and 100", 1, 100, []float64{
				100.0, 100.0 / 2, 100.0 / 3, 100.0 / 4, 100.0 / 5, 100.0 / 6, 100.0 / 7, 100.0 / 8,
				100.0 / 9, 100.0 / 10,
				100.0 / 20, 100.0 / 30, 100.0 / 40, 100.0 / 50, 100.0 / 60, 100.0 / 70, 100.0 / 80,
				100.0 / 90, 100.0 / 100,
			},
		},
		{
			"Hourly", 3600, yr,
			[]float64{
				yr, yr / 2, yr / 3, yr / 4, yr / 5, yr / 6, yr / 7, yr / 8, yr / 9, yr / 10,
				yr / 20, yr / 30, yr / 40, yr / 50, yr / 60, yr / 70, yr / 80, yr / 90, yr / 100,
				yr / 200, yr / 300, yr / 400, yr / 500, yr / 600, yr / 700, yr / 800, yr / 900, yr / 1000,
			},
		},
		{
			"Daily", 86400, yr,
			[]float64{
				yr, yr / 2, yr / 3, yr / 4, yr / 5, yr / 6, yr / 7, yr / 8, yr / 9, yr / 10,
				yr / 20, yr / 30, yr / 40, yr / 50, yr / 60, yr / 70, yr / 80, yr / 90, yr / 100,
			},
		},
		{
			"Weekly", 604800, yr,
			[]float64{
				yr, yr / 2, yr / 3, yr / 4, yr / 5, yr / 6, yr / 7, yr / 8, yr / 9, yr / 10,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			h := model.Harmonics(tc.min, tc.max)
			if len(h) != len(tc.result) {
				t.Fatalf("different lengths, expected: %v, got: %v", len(tc.result), len(h))
			}

			for i := range tc.result {
				if h[i] != tc.result[i] {
					t.Fatalf("expected %v at position %v, but got %v", tc.result[i], i, h[i])
				}
			}
		})
	}
}

func TestNewDeterministic(t *testing.T) {
	loc := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	cov := model.Diag([]float64{
		1e15, 1e5, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4,
		1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4, 1e4,
	})
	d := model.NewDeterministic(604800)

	if len(loc) != len(d.Location) {
		t.Fatalf("location length was %v, expected %v", len(d.Location), len(loc))
	}
	if len(cov) != len(d.Covariance) {
		t.Fatalf("covariance length was %v, expected %v", len(d.Covariance), len(cov))
	}

	for i := range loc {
		if loc[i] != d.Location[i] {
			t.Errorf("location expected %v at position %v, but got %v", loc[i], i, d.Location[i])
		}
	}
	for i := range cov {
		if cov[i] != d.Covariance[i] {
			t.Errorf("covariance expected %v at position %v, but got %v", cov[i], i, d.Covariance[i])
		}
	}
}

func TestDeterministicState(t *testing.T) {
	d := model.NewDeterministic(604800)

	k := d.State()

	lx, ly := k.Loc.Dims()
	if lx != 22 {
		t.Fatalf("expected location to have %v rows, but it had %v", 22, lx)
	}
	if ly != 1 {
		t.Fatalf("expected location to have %v columns, but it had %v", 1, ly)
	}

	cx, cy := k.Cov.Dims()
	if cx != 22 {
		t.Fatalf("expected covariance to have %v rows, but it had %v", 22, cx)
	}
	if cy != 22 {
		t.Fatalf("expected covariance to have %v columns, but it had %v", 22, cy)
	}
}

func TestDeterministicSystem(t *testing.T) {
	d := model.NewDeterministic(604800)

	s := d.System(100, 10, 604800)

	ax, ay := s.A.Dims()
	if ax != 22 {
		t.Fatalf("expected A to have %v rows, but it had %v", 22, ax)
	}
	if ay != 22 {
		t.Fatalf("expected A to have %v columns, but it had %v", 22, ay)
	}

	bx, by := s.B.Dims()
	if bx != 22 {
		t.Fatalf("expected B to have %v rows, but it had %v", 22, bx)
	}
	if by != 22 {
		t.Fatalf("expected B to have %v columns, but it had %v", 22, by)
	}

	cx, cy := s.C.Dims()
	if cx != 1 {
		t.Fatalf("expected B to have %v rows, but it had %v", 1, cx)
	}
	if cy != 22 {
		t.Fatalf("expected B to have %v columns, but it had %v", 22, cy)
	}

	qx, qy := s.Q.Dims()
	if qx != 22 {
		t.Fatalf("expected B to have %v rows, but it had %v", 22, qx)
	}
	if qy != 22 {
		t.Fatalf("expected B to have %v columns, but it had %v", 22, qy)
	}

	rx, ry := s.R.Dims()
	if rx != 1 {
		t.Fatalf("expected B to have %v rows, but it had %v", 1, rx)
	}
	if ry != 1 {
		t.Fatalf("expected B to have %v columns, but it had %v", 1, ry)
	}
}

func TestDeterministicUpdate(t *testing.T) {
	d := model.NewDeterministic(604800)

	resid, err := d.Update(100, 10, 604800, 1)

	if err != nil {
		t.Fatal("unexpected error during Update:", err)
	}

	if resid == 0 {
		t.Error("residual is zero")
	}

	if d.Location[0] == 0 {
		t.Error("location appears not to be updated")
	}
}

func TestDeterministicForecast(t *testing.T) {
	period := 604800.0
	n := 100

	d := model.NewDeterministic(period)
	_, err := d.Update(100, 10, period, 1)
	if err != nil {
		t.Fatal("unexpected error in Update:", err)
	}

	f := d.Forecast(period, n)

	if len(f) != n {
		t.Errorf("expected length %v, but it was %v", n, len(f))
	}
}
