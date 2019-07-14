package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplit(t *testing.T) {
	fm := make(map[int]*os.File, 128)
	f, e := os.OpenFile(filepath.Join("testdata", "1.txt"), os.O_RDONLY, 0644)
	if e != nil {
		t.Error(e)
	}
	defer f.Close()
	split(f, fm)
	if len(fm) > 128 || len(fm) < 100 {
		t.Error("expected to split to about 128 smaller files, got", len(fm))
	}

	for _, f := range fm {
		fn := f.Name()
		f.Close()
		os.Remove(fn)
	}
}
