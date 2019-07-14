package main

import (
	"bufio"
	"os"
)

func count(fm map[int]*os.File) (items []Item) {
	// count in each smaller file
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
			if (len(items) >= 100 && c >= items[0].Count) || len(items) < 100 {
				items = append(items, Item{
					Str:   s,
					Count: c,
				})
			}
		}
		topNHeapSort(items, 100)
		if len(items) > 100 {
			items = items[len(items)-1-100:] // last 100 elements
		}
	}
	return
}
