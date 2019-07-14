package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return
	}
	defer file.Close()

	fm := make(map[int]*os.File, 128)

	// split to smaller files
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		h := hash(line)
		// write line to file(h)
		f, ok := fm[h]
		if !ok {
			// create file
			f, e := os.Create(filepath.Join(tempDir, fmt.Sprintf("%d.splited.txt", h)))
			if e != nil {
				log.Fatal(e)
				return
			}
			fm[h] = f
		}
		f.WriteString(line + "\n")
	}

	// count in each smaller file
	var items []Item
	for _, f := range fm {
		f.Seek(0, 0)
		sm := make(map[string]int)

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			line := scanner.Text()
			sm[line]++
		}
		f.Close()

		for s, c := range sm {
			item := Item{
				Str:   s,
				Count: c,
			}
			items = append(items, item)
		}
		heapSort(items)
	}
}
