package main

import (
	"flag"
	"fmt"
	"log"
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
// pick up the Top N items from each smaller file
// pick up the Top N items from 128 * Top N items

func main() {
	fmt.Println("Pick up Top N items.")
	flag.StringVar(&inputFile, "input", "output.txt", "input file")
	flag.StringVar(&tempDir, "tempDir", "./", "temporary directory stores temporary files during program runs")
	flag.Parse()

	fm := make(map[int]*os.File, N)

	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return
	}

	split(file, fm)
	file.Close()

	items := count(fm)

	for _, f := range fm {
		f.Close()
	}
	for s, c := range items {
		fmt.Println(s, c)
	}
}
