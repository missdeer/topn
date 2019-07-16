package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	// N mod
	N = 128
)

var (
	inputFile string
	tempDir   string
)

// read from input file and split to 128 smaller files, hash(x) % 128
// count in each smaller file
// pick up the Top 100 items from each smaller file
// pick up the Top 100 items from (128 * Top 100) items

func main() {
	fmt.Println("Pick up Top 100 items.")
	flag.StringVar(&inputFile, "input", "output.txt", "input file")
	flag.StringVar(&tempDir, "tempDir", "./", "temporary directory stores temporary files during program runs")
	flag.Parse()

	fm := make(map[int]string, N)
	// split to about 128 smaller itnermediate file, so that it's small enough to fit the memory usage
	split(inputFile, fm)

	// count in each smaller intermediate files, and merge the sorted result
	items := count(fm, 100)

	// remove splitted smaller intermediate files
	for _, fn := range fm {
		os.Remove(fn)
	}
	// reversing
	for i := len(items)/2 - 1; i >= 0; i-- {
		opp := len(items) - 1 - i
		items[i], items[opp] = items[opp], items[i]
	}
	// output to console
	for _, c := range items {
		fmt.Println(c.Count, c.Str)
	}
}
