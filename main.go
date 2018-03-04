package main

import (
	"encoding/json"
	"fmt"

	"github.com/chulabs/seer/stream"
	"github.com/vmihailenco/msgpack"
)

func main() {
	r := &stream.Stream{}

	s, _ := stream.New("myStream", 86400, 0, 0, 1)
	b, _ := msgpack.Marshal(s)

	msgpack.Unmarshal(b, r)
	b, _ = json.Marshal(r)
	fmt.Println(string(b))
}
