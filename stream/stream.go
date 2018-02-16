package stream

import (
	"fmt"
	"time"

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
func (s *Stream) Forecast(n int, p []float64) (t []time.Time, v []float64, i []*Interval) {
	_ = s.Model.Forecast(s.Config.Period, n)
	// transform f to appropriate distribution
	// generate confidence intervals (which depends on a common dist quantiler interface)
	return
}
