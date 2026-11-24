package main

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	got := BFS(5, [][2]int{{0, 1}, {0, 2}, {1, 3}})
	want := []int{0, 1, 2, 3, 4} // level order from 0, then isolated 4
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BFS = %v, want %v", got, want)
	}

	// Cycle visited once.
	got = BFS(3, [][2]int{{0, 1}, {1, 2}, {2, 0}})
	if !reflect.DeepEqual(got, []int{0, 1, 2}) {
		t.Errorf("cycle BFS = %v", got)
	}

	if got := BFS(3, nil); !reflect.DeepEqual(got, []int{0, 1, 2}) {
		t.Errorf("edgeless = %v", got)
	}
	if got := BFS(0, nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("empty = %v", got)
	}
}
