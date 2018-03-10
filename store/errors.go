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

package store

import (
	"fmt"
)

// NotFoundError is returned when the requested Entity is not in the data store.
type NotFoundError struct {
	Kind   string
	Entity string
}

// Error message for NotFoundErrors, implements error interface.
func (err *NotFoundError) Error() string {
	return fmt.Sprintf("no %v with name %v was found in store", err.Kind, err.Entity)
}

// AlreadyExistsError is returned when an Entity name is already taken.
type AlreadyExistsError struct {
	Kind   string
	Entity string
}

// Error message for AlreadyExistsError, implements error interface.
func (err *AlreadyExistsError) Error() string {
	return fmt.Sprintf("a %v with name %v already exists", err.Kind, err.Entity)
}

// NoneFoundError is returned when a list request on a Kind returns no entities.
type NoneFoundError struct {
	Kind string
}

// Error message for NoneFoundError, implements error interface.
func (err *NoneFoundError) Error() string {
	return fmt.Sprintf("no entities of kind %v were found", err.Kind)
}

// CorruptDataError is returned when data at a key has invalid schema.
type CorruptDataError struct {
	Kind string
}

// Error message for CorruptDataError, implements error interface.
func (err *CorruptDataError) Error() string {
	return fmt.Sprintf("unable to unmarshal entity of kind %v", err.Kind)
}
