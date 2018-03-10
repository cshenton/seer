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

package kalman

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Predict predicts the next state distribution given the previous state and
// linear system equation `x_next = A * x_prev + B * w_prev `.
func Predict(p *State, m *System) (n *State, err error) {
	var loc mat.Dense
	var cov mat.Dense
	var addcov mat.Dense

	pDim := p.Dim()
	mDim, _ := m.Dims()

	if pDim != mDim {
		err = fmt.Errorf("Prev state dim must match process dim, but were %v, and %v", pDim, mDim)
		return n, err
	}

	loc.Mul(m.A, p.Loc)
	cov.Product(m.A, p.Cov, m.A.T())
	addcov.Product(m.B, m.Q, m.B.T())
	cov.Add(&cov, &addcov)
	n, err = NewState(&loc, &cov)
	return n, err
}

// Observe determines the observation distribution for the current time step
// using the linear measurement equation: `y_next = C * x_next + v_next`. It
// returns an error if the state and system process dims do not match.
func Observe(s *State, m *System) (o *State, err error) {
	var loc mat.Dense
	var cov mat.Dense

	mDim, _ := m.Dims()
	sDim := s.Dim()
	if sDim != mDim {
		err = fmt.Errorf("Prev state dim must match process dim, but were %v, and %v", sDim, mDim)
		return o, err
	}
	loc.Mul(m.C, s.Loc)
	cov.Product(m.C, s.Cov, m.C.T())
	cov.Add(&cov, m.R)
	o, err = NewState(&loc, &cov)
	return o, err
}

// StateObserve returns the observation distribution
func StateObserve(s *State, m *System) (o *State, err error) {
	var loc mat.Dense
	var cov mat.Dense

	mDim, _ := m.Dims()
	sDim := s.Dim()
	if sDim != mDim {
		err = fmt.Errorf("Prev state dim must match process dim, but were %v, and %v", sDim, mDim)
		return o, err
	}
	loc.Mul(m.C, s.Loc)
	cov.Product(m.C, s.Cov, m.C.T())
	o, err = NewState(&loc, &cov)
	return o, err
}

// Update implements the Kalman Filter update, and returns the a posteriori
// internal state distribution given the state prior and an observation.
// Also returns the post-fit residual.
func Update(p *State, m *System, v float64) (n *State, res float64, err error) {
	var (
		ir    mat.Dense
		ic    mat.Dense
		ici   mat.Dense
		gain  mat.Dense
		loc   mat.Dense
		cov   mat.Dense
		resid mat.Dense
	)
	pDim := p.Dim()
	mDim, _ := m.Dims()

	if pDim != mDim {
		err = fmt.Errorf("State dim must match process dim, but were %v, and %v", pDim, mDim)
		return n, res, err
	}

	ir.Mul(m.C, p.Loc)
	ir.Scale(-1.0, &ir)
	ir.Apply(func(i, j int, val float64) float64 { return val + v }, &ir)
	ic.Product(m.C, p.Cov, m.C.T())
	ic.Add(&ic, m.R)
	ici.Inverse(&ic)
	gain.Product(p.Cov, m.C.T(), &ici)
	loc.Mul(&gain, &ir)
	loc.Add(p.Loc, &loc)
	cov.Product(&gain, &ic, gain.T())
	cov.Scale(-1.0, &cov)
	cov.Add(p.Cov, &cov)
	resid.Mul(m.C, &loc)
	res = v - resid.At(0, 0)
	n, _ = NewState(&loc, &cov)
	return n, res, err
}
