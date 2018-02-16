package stream

import (
	"time"

	"github.com/chulabs/seer/model"
)

// Stream represents a time series data stream that can learn and forecast.
type Stream struct {
	Config *Config
	Model  *model.Model
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

// Update updates the provided sequence of values against the stream model.
func (s *Stream) Update(vals []float64) {
	for _, v := range vals {
		s.Model.Update(s.Config.Period, v)
	}
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
