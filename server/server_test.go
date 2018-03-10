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

package server_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cshenton/seer/server"
)

func testPath(t *testing.T) string {
	f, err := ioutil.TempFile(os.TempDir(), "bolt_test")
	if err != nil {
		t.Fatal("failed to create test db file")
	}
	return f.Name()
}

func TestNew(t *testing.T) {
	_, err := server.New(testPath(t))
	if err != nil {
		t.Fatal("unexpected error in server.New:", err)
	}
}

func TestNewErrs(t *testing.T) {
	_, err := server.New("/$$$NOPE!!")
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}
