package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCount(t *testing.T) {
	fm := make(map[int]*os.File, 10)
	items := count(fm)
	for s, c := range items {
		fmt.Println(s, c)
	}
}
