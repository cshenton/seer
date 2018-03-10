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

package bolt_test

import (
	"testing"
	"time"

	blt "github.com/boltdb/bolt"
	"github.com/cshenton/seer/store/bolt"
	"github.com/cshenton/seer/stream"
	"github.com/vmihailenco/msgpack"
)

func setUp(t *testing.T) (b *bolt.Store) {
	b, err := bolt.New(testPath(t))
	if err != nil {
		t.Fatal("unexpected error in bolt.New:", err)
	}
	names := []string{"sales", "visits", "usage"}

	for _, n := range names {
		s, _ := stream.New(n, 3600, 0, 0, 0)
		err := b.CreateStream(n, s)
		if err != nil {
			t.Fatal("unexpected error in CreateStream:", err)
		}
	}

	return b
}

func TestCreateStreamErrs(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	s, _ := stream.New("sales", 3600, 0, 0, 0)
	err := b.CreateStream("sales", s)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}

func TestGetStream(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	tt := []string{"sales", "visits", "usage"}

	for _, name := range tt {
		t.Run(name, func(t *testing.T) {
			s, err := b.GetStream(name)
			if err != nil {
				t.Error("unexpected error in GetStream:", err)
			}
			if s.Config.Name != name {
				t.Errorf("expected stream name %v, but got %v", name, s.Config.Name)
			}
		})
	}
}

func TestGetStreamErrs(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	s, err := b.GetStream("notastream")
	if err == nil {
		t.Error("expected error, but it was nil")
	}
	if s != nil {
		t.Error("expected nil stream, but it was", s)
	}
}

func TestDeleteStream(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	tt := []string{"sales", "visits", "usage"}

	for _, name := range tt {
		t.Run(name, func(t *testing.T) {
			err := b.DeleteStream(name)
			if err != nil {
				t.Error("unexpected error in DeleteStream:", err)
			}

			_, err = b.GetStream(name)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
		})
	}
}

func TestDeleteStreamErrs(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	err := b.DeleteStream("notastream")
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}

func TestUpdateStream(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	tt := []string{"sales", "visits", "usage"}

	for _, name := range tt {
		t.Run(name, func(t *testing.T) {
			s, err := b.GetStream(name)
			if err != nil {
				t.Error("unexpected error in GetStream:", err)
			}
			s.Update([]float64{3.14}, []time.Time{time.Now()})

			err = b.UpdateStream(name, s)
			if err != nil {
				t.Error("unexpected error in UpdateStream:", err)
			}
		})
	}
}

func TestUpdateStreamErrs(t *testing.T) {
	b := setUp(t)
	defer b.Close()
	name := "notastream"

	s, _ := stream.New(name, 3600, 0, 0, 0)
	err := b.UpdateStream(name, s)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}

func TestListStreams(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	num := 1
	size := 2

	s, err := b.ListStreams(num, size)
	if err != nil {
		t.Fatal("unexpected error in ListStreams:", err)
	}

	if len(s) != size {
		t.Errorf("expected %v streams, but there were %v", size, len(s))
	}
	for i := range s {
		if s[i].Config.Period != 3600 {
			t.Errorf("expected period of %v, but it was %v", 3600, s[i].Config.Period)
		}
	}
}

func TestListStreamsErrs(t *testing.T) {
	b := setUp(t)
	defer b.Close()

	_, err := b.ListStreams(20, 10)
	if err == nil {
		t.Error("expected error, but it was nil")
	}

	err = b.Update(func(tx *blt.Tx) error {
		bk := tx.Bucket([]byte("streams"))
		data, err := msgpack.Marshal([]string{"corrupt", "data"})
		if err != nil {
			return err
		}
		return bk.Put([]byte("corrupt"), data)
	})
	if err != nil {
		t.Fatal("unexpected error while creating corrupt data")
	}

	_, err = b.ListStreams(1, 5)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}
