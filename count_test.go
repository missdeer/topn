package main

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	fm := map[int]string{
		0: filepath.Join("testdata", "0.splited.txt"),
		1: filepath.Join("testdata", "1.splited.txt"),
		2: filepath.Join("testdata", "2.splited.txt"),
		3: filepath.Join("testdata", "3.splited.txt"),
		4: filepath.Join("testdata", "4.splited.txt"),
		5: filepath.Join("testdata", "5.splited.txt"),
		6: filepath.Join("testdata", "6.splited.txt"),
		7: filepath.Join("testdata", "7.splited.txt"),
		8: filepath.Join("testdata", "8.splited.txt"),
		9: filepath.Join("testdata", "9.splited.txt"),
	}
	items := count(fm, 3)
	if len(items) != 3 {
		t.Error("len(items) should be 3")
	}
	expected1 := []Item{
		{"2", 11},
		{"3", 12},
		{"5", 17},
	}
	if !reflect.DeepEqual(items, expected1) {
		t.Errorf("expected %v, got %v", expected1, items)
	}
}
