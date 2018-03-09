package bolt_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/chulabs/seer/store/bolt"
)

func testPath(t *testing.T) string {
	f, err := ioutil.TempFile(os.TempDir(), "bolt_test")
	if err != nil {
		t.Fatal("failed to create test db file")
	}
	return f.Name()
}

func TestNew(t *testing.T) {
	b, err := bolt.New(testPath(t))
	defer b.Close()
	if err != nil {
		t.Fatal("unexpected error in bolt.New:", err)
	}
}

func TestNewErrs(t *testing.T) {
	_, err := bolt.New("/$$$NOPE!!")
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}
