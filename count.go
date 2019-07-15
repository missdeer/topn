package main

import (
	"bufio"
	"os"
)

func count(fm map[int]*os.File, topN int) (items []Item) {
	// count in each smaller file
	orderred := false
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
			if !orderred ||
				(orderred && len(items) >= topN && c >= items[0].Count) || // drop smaller ones
				len(items) < topN {
				items = append(items, Item{
					Str:   s,
					Count: c,
				})
			}
		}

		// get the top N items
		topNHeapSort(items, topN)
		if len(items) >= topN {
			orderred = true
			items = items[len(items)-topN:] // keep the last N elements
		}
	}
	return
}
