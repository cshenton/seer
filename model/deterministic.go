package model

import "github.com/chulabs/seer/dist/mv"

// Deterministic is the type against which we apply deterministic model updates.
type Deterministic struct {
	*mv.Normal
}
