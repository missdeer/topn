package main

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []Item{
		{"a", 12},
		{"b", 11},
		{"c", 13},
		{"d", 5},
		{"e", 6},
		{"f", 7},
	}
	expected := []Item{
		{"d", 5},
		{"e", 6},
		{"f", 7},
		{"b", 11},
		{"a", 12},
		{"c", 13},
	}
	heapSort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}
