/*
 * Copyright (C) 2018 The Seer Authors. All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
