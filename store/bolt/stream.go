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
	"github.com/cshenton/seer/store"
	"github.com/cshenton/seer/stream"
	"github.com/vmihailenco/msgpack"

	// Avoid namespace conflicts
	blt "github.com/boltdb/bolt"
)

// streamBucket is the key for the stream bucket.
var streamBucket = []byte("streams")

// streamInit idempotently sets up the store to be ready to store streams.
func (b *Store) streamInit() {
	b.Update(func(tx *blt.Tx) error {
		tx.CreateBucketIfNotExists(streamBucket)
		return nil
	})
}

// CreateStream saves the provided stream at name, returns an error if a
// stream already exists at that address.
func (b *Store) CreateStream(name string, s *stream.Stream) (err error) {
	err = b.Update(func(tx *blt.Tx) error {
		bk := tx.Bucket(streamBucket)

		val := bk.Get([]byte(name))
		if val != nil {
			return &store.AlreadyExistsError{Kind: "stream", Entity: name}
		}

		val, _ = msgpack.Marshal(s)
		err := bk.Put([]byte(name), val)
		return err
	})

	return err
}

// GetStream returns the stream stored at name, or an error if the stream does
// not exist, or has corrupted data.
func (b *Store) GetStream(name string) (s *stream.Stream, err error) {
	s = &stream.Stream{}

	err = b.View(func(tx *blt.Tx) error {
		bk := tx.Bucket(streamBucket)

		val := bk.Get([]byte(name))
		if val == nil {
			return &store.NotFoundError{Kind: "stream", Entity: name}
		}
		err = msgpack.Unmarshal(val, s)
		return err
	})

	if err != nil {
		return nil, err
	}
	return s, nil
}

// DeleteStream deletes the stream stored at name, or returns an error if
// no such stream exists.
func (b *Store) DeleteStream(name string) (err error) {
	err = b.Update(func(tx *blt.Tx) error {
		bk := tx.Bucket(streamBucket)

		val := bk.Get([]byte(name))
		if val == nil {
			return &store.NotFoundError{Kind: "stream", Entity: name}
		}

		err := bk.Delete([]byte(name))
		return err
	})

	return err
}

// UpdateStream overwrites the stream at name with the provided stream, or
// returns an error if no stream exists at name.
func (b *Store) UpdateStream(name string, s *stream.Stream) (err error) {
	err = b.Update(func(tx *blt.Tx) error {
		bk := tx.Bucket(streamBucket)

		val := bk.Get([]byte(name))
		if val == nil {
			return &store.NotFoundError{Kind: "stream", Entity: name}
		}

		val, _ = msgpack.Marshal(s)
		err := bk.Put([]byte(name), val)
		return err
	})

	return err
}

// ListStreams returns a paged list of streams, or an error if none are found.
func (b *Store) ListStreams(pageNum, pageSize int) (s []*stream.Stream, err error) {
	err = b.View(func(tx *blt.Tx) error {
		bk := tx.Bucket(streamBucket)
		offset := (pageNum - 1) * pageSize

		c := bk.Cursor()
		i := 0

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if i >= offset+pageSize {
				break
			}
			if i >= offset {
				st := &stream.Stream{}
				err = msgpack.Unmarshal(v, st)
				if err != nil {
					return err
				}
				s = append(s, st)
			}
			i++
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		err = &store.NoneFoundError{Kind: "stream"}
		return nil, err
	}
	return s, nil
}
