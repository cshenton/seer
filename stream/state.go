package stream

import (
	"time"
)

// History contains the last two observations for this stream, it is required
// to do covariance estimation in real time.
type History [2]float64

// Update shifts the history one place down and adds a new value.
func (h History) Update(v float64) {
	h[1] = h[0]
	h[0] = v
}

// State stores dynamic state about a stream.
type State struct {
	Time          time.Time
	Deterministic Something
	Stochastic    Something
	Theta         float64
	Zeta          float64
	History       History
}

// NewState initialises a state given a stream config.
func NewState(c *Config) (s *State, err error) {
	s = &State{}
	return
}
