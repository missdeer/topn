package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	fm := make(map[int]*os.File, 10)
	fm[0], _ = os.OpenFile(filepath.Join("testdata", "0.splited.txt"), os.O_RDONLY, 0644)
	fm[1], _ = os.OpenFile(filepath.Join("testdata", "1.splited.txt"), os.O_RDONLY, 0644)
	fm[2], _ = os.OpenFile(filepath.Join("testdata", "2.splited.txt"), os.O_RDONLY, 0644)
	fm[3], _ = os.OpenFile(filepath.Join("testdata", "3.splited.txt"), os.O_RDONLY, 0644)
	fm[4], _ = os.OpenFile(filepath.Join("testdata", "4.splited.txt"), os.O_RDONLY, 0644)
	fm[5], _ = os.OpenFile(filepath.Join("testdata", "5.splited.txt"), os.O_RDONLY, 0644)
	fm[6], _ = os.OpenFile(filepath.Join("testdata", "6.splited.txt"), os.O_RDONLY, 0644)
	fm[7], _ = os.OpenFile(filepath.Join("testdata", "7.splited.txt"), os.O_RDONLY, 0644)
	fm[8], _ = os.OpenFile(filepath.Join("testdata", "8.splited.txt"), os.O_RDONLY, 0644)
	fm[9], _ = os.OpenFile(filepath.Join("testdata", "9.splited.txt"), os.O_RDONLY, 0644)
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
