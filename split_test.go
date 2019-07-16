package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplit(t *testing.T) {
	fm := make(map[int]string, 128)
	split(filepath.Join("testdata", "1.txt"), fm)
	if len(fm) > 128 || len(fm) < 100 {
		t.Error("expected to split to about 128 smaller files, got", len(fm))
	}

	for _, fn := range fm {
		os.Remove(fn)
	}
}

func TestSplit2(t *testing.T) {
	fm := make(map[int]string, 128)
	split2(filepath.Join("testdata", "1.txt"), fm)
	if len(fm) > 128 || len(fm) < 100 {
		t.Error("expected to split to about 128 smaller files, got", len(fm))
	}

	for _, fn := range fm {
		os.Remove(fn)
	}
}
