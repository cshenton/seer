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

package store_test

import (
	"testing"

	"github.com/cshenton/seer/store"
)

func TestNotFoundError(t *testing.T) {
	msg := "no stream with name wallace was found in store"
	err := store.NotFoundError{
		Kind:   "stream",
		Entity: "wallace",
	}

	if err.Error() != msg {
		t.Errorf("expected message `%v`, but got `%v`", msg, err.Error())
	}
}

func TestAlreadyExistsError(t *testing.T) {
	msg := "a stream with name wallace already exists"
	err := store.AlreadyExistsError{
		Kind:   "stream",
		Entity: "wallace",
	}

	if err.Error() != msg {
		t.Errorf("expected message `%v`, but got `%v`", msg, err.Error())
	}
}

func TestNoneFoundError(t *testing.T) {
	msg := "no entities of kind stream were found"
	err := store.NoneFoundError{
		Kind: "stream",
	}

	if err.Error() != msg {
		t.Errorf("expected message `%v`, but got `%v`", msg, err.Error())
	}
}

func TestCorruptDataError(t *testing.T) {
	msg := "unable to unmarshal entity of kind stream"
	err := store.CorruptDataError{
		Kind: "stream",
	}

	if err.Error() != msg {
		t.Errorf("expected message `%v`, but got `%v`", msg, err.Error())
	}
}
