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

package uv_test

import (
	"math"
	"testing"

	"github.com/cshenton/seer/dist/uv"
)

func TestConfidenceInterval(t *testing.T) {
	n, _ := uv.NewNormal(0, 1)
	l, u, err := uv.ConfidenceInterval(n, 0.9)
	if err != nil {
		t.Fatal("unexpected error in ConfidenceInterval,", err)
	}
	if math.Abs(l - -1.645) > 1e-3 {
		t.Errorf("expected lower bound %v, but got %v", -1.645, l)
	}
	if math.Abs(u-1.645) > 1e-3 {
		t.Errorf("expected upper bound %v, but got %v", 1.645, u)
	}

	_, _, err = uv.ConfidenceInterval(n, -1)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}
