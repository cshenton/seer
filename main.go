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

package main

import (
	"log"
	"net"
	"path/filepath"

	"github.com/cshenton/seer/seer"
	"github.com/cshenton/seer/server"
	"google.golang.org/grpc"
)

var port = ":8080"
var path = filepath.FromSlash("/var/seer")

func main() {
	srv, err := server.New(path)
	if err != nil {
		log.Fatal("failed to create server:", err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	s := grpc.NewServer()
	seer.RegisterSeerServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
