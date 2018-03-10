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
