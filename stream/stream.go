package stream

import "github.com/chulabs/seer/model"

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
