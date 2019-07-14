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
func TestTopNHeapSort(t *testing.T) {
	arr := []Item{
		{"a", 12},
		{"b", 11},
		{"c", 13},
		{"d", 5},
		{"e", 6},
		{"f", 7},
	}
	expected3 := []Item{
		{"b", 11},
		{"a", 12},
		{"c", 13},
	}
	topNHeapSort(arr, 3)
	if !reflect.DeepEqual(arr[3:], expected3) {
		t.Errorf("expected %v, got %v", expected3, arr)
	}

	arr = []Item{
		{"c", 13},
		{"a", 12},
		{"b", 11},
		{"f", 7},
		{"e", 6},
		{"d", 5},
	}
	expected2 := []Item{
		{"a", 12},
		{"c", 13},
	}
	topNHeapSort(arr, 2)
	if !reflect.DeepEqual(arr[4:], expected2) {
		t.Errorf("expected %v, got %v", expected2, arr)
	}

	arr = []Item{
		{"c", 13},
		{"a", 12},
		{"b", 11},
		{"d", 5},
		{"e", 6},
		{"f", 7},
	}
	expected1 := []Item{
		{"c", 13},
	}
	topNHeapSort(arr, 1)
	if !reflect.DeepEqual(arr[5:], expected1) {
		t.Errorf("expected %v, got %v", expected1, arr)
	}

	arr = []Item{
		{"c", 13},
		{"a", 12},
		{"b", 11},
		{"d", 5},
		{"e", 6},
		{"f", 7},
	}
	expected4 := []Item{
		{"d", 5},
		{"e", 6},
		{"f", 7},
		{"b", 11},
		{"a", 12},
		{"c", 13},
	}
	topNHeapSort(arr, 100)
	if !reflect.DeepEqual(arr, expected4) {
		t.Errorf("expected %v, got %v", expected4, arr)
	}
}
