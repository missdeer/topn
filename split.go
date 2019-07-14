package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func split(file *os.File, fm map[int]*os.File) {
	// split to smaller files
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
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
}
