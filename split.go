package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func split2(inputFile string, fm map[int]string) {
	file, e := os.OpenFile(inputFile, os.O_RDONLY, 0644)
	if e != nil {
		log.Fatal(e)
		return
	}
	defer file.Close()

	// split to smaller files
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		h := hash2(line) + N
		// write line to file(h)
		var f *os.File
		var e error
		if fn, ok := fm[h]; !ok {
			// create file
			fn = fmt.Sprintf("%s.%d.splited.txt", inputFile[:len(inputFile)-len(".splited.txt")], h)
			f, e = os.Create(fn)
			if e != nil {
				log.Fatal(e)
				return
			}
			fm[h] = fn
		} else {
			// open file to append
			f, e = os.OpenFile(fn, os.O_APPEND, 0644)
			if e != nil {
				log.Fatal(e)
				return
			}
		}

		f.WriteString(line + "\n")
		f.Close()
	}
}

func split(inputFile string, fm map[int]string) {
	file, e := os.OpenFile(inputFile, os.O_RDONLY, 0644)
	if e != nil {
		log.Fatal(e)
		return
	}
	defer file.Close()

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
		var f *os.File
		var e error
		if fn, ok := fm[h]; !ok {
			// create file
			fn = filepath.Join(tempDir, fmt.Sprintf("%d.splited.txt", h))
			f, e = os.Create(fn)
			if e != nil {
				log.Fatal(e)
				return
			}
			fm[h] = fn
		} else {
			// open file to append
			f, e = os.OpenFile(fn, os.O_APPEND, 0644)
			if e != nil {
				log.Fatal(e)
				return
			}
		}

		f.WriteString(line + "\n")
		f.Close()
	}

	for _, fn := range fm {
		fi, e := os.Stat(fn)
		if e != nil {
			log.Fatal(e)
			return
		}
		if fi.Size() > 1000000000 {
			// split again
			split(fn, fm)
			os.Remove(fn)
		}
	}
}
