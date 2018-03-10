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

package bolt

import (
	// Avoid namespace conflicts
	blt "github.com/boltdb/bolt"
)

// Store wraps a bolt DB and fulfills the store.StreamStore interface.
type Store struct {
	*blt.DB
}

// New creates a bolt Store at the provided file path.
func New(path string) (b *Store, err error) {
	db, err := blt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	b = &Store{db}

	b.streamInit()

	return b, nil
}
