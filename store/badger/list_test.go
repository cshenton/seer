package badger_test

import (
	"testing"

	"github.com/chulabs/seer/store/badger"
)

func TestStreamListContains(t *testing.T) {
	tt := []struct {
		name     string
		list     []string
		element  string
		contains bool
	}{
		{"does contain", []string{"apple", "square"}, "apple", true},
		{"does not contain", []string{"apple", "square"}, "circle", false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sl := badger.StreamList(tc.list)

			c := sl.Contains(tc.element)

			if c != tc.contains {
				t.Errorf("expected Contains to return %v, but it got %v", tc.contains, c)
			}
		})
	}
}

func TestStreamListRemove(t *testing.T) {
	tt := []struct {
		name    string
		list    []string
		element string
		result  []string
	}{
		{"element in middle", []string{"apple", "square", "three"}, "square", []string{"apple", "three"}},
		{"element at end", []string{"apple", "square", "three"}, "three", []string{"apple", "square"}},
		{"does not contain", []string{"apple", "square"}, "one", []string{"apple", "square"}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sl := badger.StreamList(tc.list)
			res := sl.Remove(tc.element)

			for i := range tc.result {
				if tc.result[i] != res[i] {
					t.Errorf("expected %v at position %v, but got %v", tc.result[i], i, res[i])
				}
			}
		})
	}

}

func TestStreamListAdd(t *testing.T) {
	result := badger.StreamList([]string{"a", "b", "c", "f"})

	sl := badger.StreamList([]string{"a", "b", "c"})
	res := sl.Add("f")

	for i := range result {
		if result[i] != res[i] {
			t.Errorf("expected %v at position %v, but got %v", result[i], i, res[i])
		}
	}
}
