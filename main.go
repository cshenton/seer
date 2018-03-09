package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func testPath(t *testing.T) string {
	path, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal("failed to construct test path")
	}
	return path
}

func main() {
	f, _ := ioutil.TempFile(os.TempDir(), "bolt_test")
	fmt.Println(f.Name())
}
