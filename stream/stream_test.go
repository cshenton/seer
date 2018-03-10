package stream_test

import (
	"testing"
	"time"

	"github.com/cshenton/seer/stream"
)

func TestNew(t *testing.T) {
	s, err := stream.New("sales", 60, 10, 0, 1)

	if err != nil {
		t.Error("unexpected error in New:", err)
	}
	if s.Config.Min != 10 {
		t.Errorf("expected config min %v, but got %v", 10, s.Config.Min)
	}
}

func TestNewErrs(t *testing.T) {
	s, err := stream.New("y", 60, 0, 0, 0)

	if err == nil {
		t.Error("expected error but got nil")
	}
	if s != nil {
		t.Error("expected nil stream but got:", s)
	}
}

func TestStreamUpdate(t *testing.T) {
	vals := []float64{1, 2}
	times := []time.Time{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC)}

	s, err := stream.New("streamy", 86400, 0, 0, 0)
	if err != nil {
		t.Fatal("unexpected error in New:", err)
	}

	err = s.Update(vals, times)
	if err != nil {
		t.Fatal("unexpected error in Update:", err)
	}

	if s.Model.RCE.History[0] == 0 {
		t.Error("model was not updated")
	}
}

func TestStreamUpdateErrs(t *testing.T) {
	tt := []struct {
		name   string
		times  []time.Time
		values []float64
	}{
		{"mismatched lengths", []time.Time{time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC)}, []float64{1, 2}},
		{"wrong next time", []time.Time{time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC)}, []float64{1}},
		{
			"wrong intermediate time",
			[]time.Time{
				time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2016, 1, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2016, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			[]float64{2, 1, 3},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := stream.New("streamy", 86400, 0, 0, 0)
			if err != nil {
				t.Fatal("unexpected error in New:", err)
			}
			err = s.Update([]float64{1}, []time.Time{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)})
			if err != nil {
				t.Fatal("unexpected error in Update:", err)
			}

			err = s.Update(tc.values, tc.times)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
		})
	}
}

func TestStreamForecast(t *testing.T) {
	tt := []struct {
		name   string
		n      int
		probs  []float64
		domain int
	}{
		{"single period, single prob", 1, []float64{0.9}, 0},
		{"multiple periods, single prob", 10, []float64{0.9}, 0},
		{"single period, multiple probs", 1, []float64{0.9, 0.99}, 0},
		{"multiple periods, multiple probs", 10, []float64{0.9, 0.99}, 0},
		{"single period, single prob, right continuous", 1, []float64{0.9}, 1},
		{"multiple periods, single prob, right continuous", 10, []float64{0.9}, 1},
		{"single period, multiple probs, right continuous", 1, []float64{0.9, 0.99}, 1},
		{"multiple periods, multiple probs, right continuous", 10, []float64{0.9, 0.99}, 1},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, _ := stream.New("stream", 3600, 0, 0, tc.domain)
			s.Update([]float64{1}, []time.Time{time.Now()})

			tm, v, in, err := s.Forecast(tc.n, tc.probs)

			if err != nil {
				t.Fatal("unexpected error in Forecast,", err)
			}
			if len(tm) != tc.n {
				t.Errorf("expected %v times, but there were %v", tc.n, len(tm))
			}
			if len(v) != tc.n {
				t.Errorf("expected %v values, but there were %v", tc.n, len(v))
			}
			if len(in) != len(tc.probs) {
				t.Fatalf("expected %v intervals, but there were %v", len(tc.probs), len(in))
			}
			for i := range in {
				if in[i].Probability != tc.probs[i] {
					t.Errorf(
						"expected interval %v to have probability %v, but it was %v",
						i, tc.probs[i], in[i].Probability,
					)
				}
				if len(in[i].LowerBound) != tc.n {
					t.Errorf("expected %v lowerbounds, but there were %v", tc.n, len(in[i].LowerBound))
				}
				if len(in[i].UpperBound) != tc.n {
					t.Errorf("expected %v lowerbounds, but there were %v", tc.n, len(in[i].UpperBound))
				}
			}
		})
	}
}

func TestStreamForecastErrs(t *testing.T) {
	tt := []struct {
		name  string
		n     int
		probs []float64
	}{
		{"bad probs", 10, []float64{0.5, 2}},
		{"zero length", 0, []float64{0.9, 0.99}},
		{"negative length", -3, []float64{0.9, 0.99}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, _ := stream.New("stream", 3600, 0, 0, 0)
			s.Update([]float64{1}, []time.Time{time.Now()})

			_, _, _, err := s.Forecast(tc.n, tc.probs)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
		})
	}
}
