package model

import "github.com/chulabs/seer/dist/mv"

// Stochastic is the type against which we apply stochastic model updates.
type Stochastic struct {
	*mv.Normal
}
