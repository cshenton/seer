package stream

// Domain determines whether data are continuous or discrete.
type Domain int

// Valid values for Domain. These MUST match with the enum defined in the
// protocol buffer.
const (
	Continuous         Domain = 0
	ContinuousRight    Domain = 1
	ContinuousInterval Domain = 2
	DiscreteRight      Domain = 3
	DiscreteInterval   Domain = 4
)

// IsInterval returns whether the domain is restricted to an interval
func (d Domain) IsInterval() bool {
	return d == ContinuousInterval || d == DiscreteInterval
}
