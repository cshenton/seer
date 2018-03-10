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

package uv

import "errors"

// Quantiler defines distributions that have computable quantiles. That is,
// which accept a probability on [0,1] and return a member in their support.
type Quantiler interface {
	Quantile(p float64) (q float64, err error)
}

// ConfidenceInterval constructs a confidence interval with confidence level p,
// where higher confidence in [0,1] means a wider interval.
func ConfidenceInterval(q Quantiler, p float64) (l, u float64, err error) {
	if p < 0 || p > 1 {
		err := errors.New("probabilities must be between 0 and 1")
		return 0, 0, err
	}
	l, _ = q.Quantile(0.5 - p/2)
	u, _ = q.Quantile(0.5 + p/2)
	return l, u, nil
}
