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
	"time"
)

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

// IsInterval returns whether the domain is restricted to an interval.
func (d Domain) IsInterval() bool {
	return d == ContinuousInterval || d == DiscreteInterval
}

// IsRight returns whether the domain is open to the right only.
func (d Domain) IsRight() bool {
	return d == ContinuousRight || d == DiscreteRight
}

// IsOpen returns whether the domain is unconstrained.
func (d Domain) IsOpen() bool {
	return d == Continuous
}

// Config stores static configuration about a stream.
type Config struct {
	Name   string
	Period float64
	Min    float64
	Max    float64
	Domain Domain
}

// NewConfig validates the provided configuration data and returns a Config.
func NewConfig(name string, period, min, max float64, domain int) (c *Config, err error) {
	dom := Domain(domain)
	if dom.IsInterval() && max <= min {
		err = errors.New(`max must be greater than min for interval domain`)
		return nil, err
	}

	if len(name) < 3 {
		err = errors.New(`name must be three characters or longer`)
		return nil, err
	}

	if period < 1.0 {
		err = errors.New(`period must be 1s or longer`)
		return nil, err
	}

	c = &Config{
		Name:   name,
		Period: period,
		Min:    min,
		Max:    max,
		Domain: dom,
	}
	return c, nil
}

// Duration constructs a time.Duration from the period.
func (c *Config) Duration() time.Duration {
	return time.Duration(c.Period * 1e9)
}
