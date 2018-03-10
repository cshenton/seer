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

package server

import (
	"github.com/cshenton/seer/store"
	"github.com/cshenton/seer/store/bolt"
)

// Server fulfills the protocol buffer's SeerServer interface.
type Server struct {
	DB store.StreamStore
}

// New creates a database connection and returns a Server.
func New(path string) (srv *Server, err error) {
	db, err := bolt.New(path)
	if err != nil {
		return nil, err
	}
	return &Server{DB: db}, nil
}
