package stream

import (
	"errors"
	"math"

	"github.com/cshenton/seer/dist/uv"
)

// ToLogNormal returns a log normal distribution with the same first and
// second moments as the input normal distribution.
func ToLogNormal(n *uv.Normal) (ln *uv.LogNormal, err error) {
	if n.Location <= 0 {
		err := errors.New("Must have strictly positive location to transform to log normal")
		return nil, err
	}
	scale := math.Sqrt(math.Log1p(math.Pow(n.Scale/n.Location, 2)))
	loc := math.Log(n.Location) - math.Log1p(math.Pow(n.Scale/n.Location, 2))/2

	ln, _ = uv.NewLogNormal(loc, scale)
	return ln, nil
}
