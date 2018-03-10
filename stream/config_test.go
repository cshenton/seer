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

package stream_test

import (
	"testing"

	"github.com/chulabs/seer/stream"
)

func TestDomainIsInterval(t *testing.T) {
	tt := []struct {
		name     string
		domain   stream.Domain
		interval bool
	}{
		{"continuous", stream.Continuous, false},
		{"continuous right", stream.ContinuousRight, false},
		{"continuous interval", stream.ContinuousInterval, true},
		{"discrete right", stream.DiscreteRight, false},
		{"discrete interval", stream.DiscreteInterval, true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.domain.IsInterval()
			if tc.interval != res {
				t.Errorf("expected IsInterval to be %v, but it was %v", tc.interval, res)
			}
		})
	}
}

func TestDomainIsRight(t *testing.T) {
	tt := []struct {
		name     string
		domain   stream.Domain
		interval bool
	}{
		{"continuous", stream.Continuous, false},
		{"continuous right", stream.ContinuousRight, true},
		{"continuous interval", stream.ContinuousInterval, false},
		{"discrete right", stream.DiscreteRight, true},
		{"discrete interval", stream.DiscreteInterval, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.domain.IsRight()
			if tc.interval != res {
				t.Errorf("expected IsRight to be %v, but it was %v", tc.interval, res)
			}
		})
	}
}

func TestDomainIsOpen(t *testing.T) {
	tt := []struct {
		name     string
		domain   stream.Domain
		interval bool
	}{
		{"continuous", stream.Continuous, true},
		{"continuous right", stream.ContinuousRight, false},
		{"continuous interval", stream.ContinuousInterval, false},
		{"discrete right", stream.DiscreteRight, false},
		{"discrete interval", stream.DiscreteInterval, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.domain.IsOpen()
			if tc.interval != res {
				t.Errorf("expected IsOpen to be %v, but it was %v", tc.interval, res)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	tt := []struct {
		name   string
		period float64
		min    float64
		max    float64
		domain int
	}{
		{"Invalid but redundant bounds", 60, 100, 50, 0},
		{"Normal Config", 60, 0, 0, 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c, err := stream.NewConfig(tc.name, tc.period, tc.min, tc.max, tc.domain)
			if err != nil {
				t.Error("unexpected error in NewConfig:", err)
			}
			if c.Name != tc.name {
				t.Errorf("expected name %v, but got %v", tc.name, c.Name)
			}
		})
	}
}

func TestNewConfigErrs(t *testing.T) {
	tt := []struct {
		name   string
		period float64
		min    float64
		max    float64
		domain int
	}{
		{"Max less than min", 60, 100, 50, 2},
		{"Nm", 60, 0, 0, 0},
		{"Short period", 0.01, 0, 100, 4},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c, err := stream.NewConfig(tc.name, tc.period, tc.min, tc.max, tc.domain)
			if err == nil {
				t.Error("expected error, but got nil")
			}
			if c != nil {
				t.Error("expected nil pointer, but got", c)
			}
		})
	}
}
