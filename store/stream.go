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
	"context"

	"github.com/cshenton/seer/stream"
)

// StreamStore defines the methods required to store and retrieve streams.
type StreamStore interface {
	CreateStream(name string, s *stream.Stream) (err error)
	GetStream(name string) (s *stream.Stream, err error)
	DeleteStream(name string) (err error)
	ListStreams(pageNum, pageSize int) (s []*stream.Stream, err error)
	UpdateStream(name string, s *stream.Stream) (err error)
}

// CreateStream creates a stream using the store on the current context, it returns an
// error if the stream already exists.
func CreateStream(c context.Context, name string, s *stream.Stream) (err error) {
	return streamFromContext(c).CreateStream(name, s)
}

// GetStream returns the stream with the specific name using the current context store.
func GetStream(c context.Context, name string) (s *stream.Stream, err error) {
	return streamFromContext(c).GetStream(name)
}

// DeleteStream deletes the stream with the specific name using the current context store.
func DeleteStream(c context.Context, name string) (err error) {
	return streamFromContext(c).DeleteStream(name)
}

// ListStreams lists pages streams from the current context store.
func ListStreams(c context.Context, pageNum, pageSize int) (s []*stream.Stream, err error) {
	return streamFromContext(c).ListStreams(pageNum, pageSize)
}

// UpdateStream saves the provided stream, and returns an error if no stream with
// the given name exists.
func UpdateStream(c context.Context, name string, s *stream.Stream) (err error) {
	return streamFromContext(c).UpdateStream(name, s)
}
