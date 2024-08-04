package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTopologicalSort(t *testing.T) {
	t.Run("Order in normal case", func(t *testing.T) {
		commits := map[Hash]Commit{
			"a": {Hash: "a", Children: []Hash{"b", "c"}, AuthorDate: time.Now()},
			"b": {Hash: "b", Children: []Hash{"d"}, AuthorDate: time.Now()},
			"c": {Hash: "c", Children: []Hash{"e"}, AuthorDate: time.Now()},
			"d": {Hash: "d", Children: []Hash{"f"}, AuthorDate: time.Now()},
			"e": {Hash: "e", Children: []Hash{}, AuthorDate: time.Now()},
			"f": {Hash: "f", Children: []Hash{}, AuthorDate: time.Now()},
		}
        newests := []Hash{"f", "e", "d", "c", "b", "a"}

		// the graph is:
		// a -> b -> d -> f
		//   ->  c ->  e
		got := TopologicalSort(commits, newests)
		want := []Hash{"f", "e", "d", "c", "b", "a"}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Order in case children older parent", func(t *testing.T) {
		// commits was in order of AuthorDate
		commits := map[Hash]Commit{
			"a": {Hash: "a", Children: []Hash{"b", "c"}, AuthorDate: time.Now()},
			"b": {Hash: "b", Children: []Hash{"d", "e"}, AuthorDate: time.Now()},
			"c": {Hash: "c", Children: []Hash{"f"}, AuthorDate: time.Now()},
			"d": {Hash: "d", Children: []Hash{"g"}, AuthorDate: time.Now()},
			"h": {Hash: "h", Children: []Hash{"j"}, AuthorDate: time.Now()}, // h is older than e
			"e": {Hash: "e", Children: []Hash{"h"}, AuthorDate: time.Now()},
			"f": {Hash: "f", Children: []Hash{"i"}, AuthorDate: time.Now()},
			"g": {Hash: "g", Children: []Hash{"j"}, AuthorDate: time.Now()},
			"i": {Hash: "i", Children: []Hash{"j"}, AuthorDate: time.Now()},
			"j": {Hash: "j", Children: []Hash{}, AuthorDate: time.Now()},
		}
        newests := []Hash{"j", "i", "g", "f", "e", "h", "d", "c", "b", "a"}

		// the graph is:
		// a -> b -> d -> g -> j
		//   ->        e -> (h) -> j
		//   ->   c ->   f -> i -> j
		got := TopologicalSort(commits, newests)
		want := []Hash{"j", "i", "g" , "f", "h", "e", "d", "c", "b", "a"}

        if reflect.DeepEqual(got, want) == false {
            t.Errorf("got %v, want %v", got, want)
        }
	})
}
