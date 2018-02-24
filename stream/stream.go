package stream

import (
	"fmt"
	"time"

	"github.com/chulabs/seer/dist/uv"
	"github.com/chulabs/seer/model"
)

// Stream represents a time series data stream that can learn and forecast.
type Stream struct {
	Config *Config
	Model  *model.Model
	Time   time.Time
}

// New constructs a stream given the required data.
func New(name string, period, min, max float64, domain int) (s *Stream, err error) {
	conf, err := NewConfig(name, period, min, max, domain)
	if err != nil {
		return nil, err
	}
	s = &Stream{
		Config: conf,
		Model:  model.New(conf.Period),
	}
	return s, nil
}

// Update updates the provided sequence of values against the stream model. It
// returns an error if the times are not in sequence, or if there are an
// incorrect number of corresponding values.
func (s *Stream) Update(vals []float64, times []time.Time) (err error) {
	if len(vals) != len(times) {
		err = fmt.Errorf("vals, times should be equal length, but were %v and %v", len(vals), len(times))
		return err
	}

	var t time.Time
	if s.Time.IsZero() {
		t = times[0]
	} else {
		t = s.Time.Add(time.Duration(s.Config.Period * 1e9))
	}

	for i := range times {
		if times[i] != t {
			err = fmt.Errorf("expected time %v at position %v, but got %v", t, i, times[i])
			return err
		}
		t = t.Add(time.Duration(s.Config.Period * 1e9))
	}

	for _, v := range vals {
		s.Model.Update(s.Config.Period, v)
	}
	s.Time = times[len(times)-1]
	return nil
}

// Interval is a forecast confidence interval.
type Interval struct {
	Probability float64
	LowerBound  []float64
	UpperBound  []float64
}

// Forecast forecasts against the model and transforms the result to the appropriate domain.
func (s *Stream) Forecast(n int, probs []float64) (t []time.Time, v []float64, in []*Interval) {
	f := s.Model.Forecast(s.Config.Period, n)
	q := make([]uv.Quantiler, n)

	switch s.Config.Domain {
	case Continuous:
		q = nil
	case ContinuousRight:
		q = nil
	case ContinuousInterval:
		q = nil
	case DiscreteRight:
		q = nil
	case DiscreteInterval:
		q = nil
	}

	t = make([]time.Time, n)
	v = make([]float64, n)
	in = make([]*Interval, len(probs))
	for i := range in {
		in[i].Probability = probs[i]
		in[i].LowerBound = make([]float64, n)
		in[i].UpperBound = make([]float64, n)
	}

	prev := s.Time
	for i := range t {
		next := prev.Add(s.Config.Duration())

		t[i] = next
		v[i], _ = q[i].Quantile(0.5)

		for j := range in {
			l, u, _ := uv.ConfidenceInterval(q[i], in[j].Probability)
			in[j].LowerBound[i] = l
			in[j].UpperBound[i] = u
		}

		prev = next
	}
	return t, v, in
}
