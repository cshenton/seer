package stream

import (
	"time"

	"github.com/chulabs/seer/model"
)

// State stores dynamic state about a stream.
type State struct {
	Time          time.Time
	Deterministic *model.Deterministic
	Stochastic    *model.Stochastic
	RCE           *model.RCE
}

// NewState initialises a state given a stream config.
func NewState(c *Config) (s *State, err error) {
	s = &State{
		Deterministic: nil,
		Stochastic:    model.NewStochastic(),
		RCE:           model.NewRCE(),
	}
	return s, nil
}

// Update iterates the State in response to an observed event.
func (s *State) Update(v, period float64) {
	// get covariances from RCE
	// construct filters using period (and external package)
	// forward pass deterministic filter, get residual, state
	// update rce on residual
	// forward pass stochastic filter, state
	// increment time
	// apply new time
	// apply new states
	// apply new RCE
}

// Forecast passes forward through the filter and returns
