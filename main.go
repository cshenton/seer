package main

import (
	"encoding/json"
	"fmt"

	"github.com/chulabs/seer/stream"
)

func main() {
	s, _ := stream.New("myStream", 86400, 0, 0, 1)
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
}
