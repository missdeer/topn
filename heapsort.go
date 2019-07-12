package main

// Item array element type that would be sorted by Count
type Item struct {
	Str   string
	Count int
}

func heapSort(arr []Item) {
	n := len(arr)
	for i := n; i > -1; i-- {
		heapifyItem(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapifyItem(arr, i, 0)
	}
}

func heapifyItem(arr []Item, n int, i int) {
	largest := i
	l := i*2 + 1
	r := i*2 + 2

	if l < n && arr[i].Count < arr[l].Count {
		largest = l
	}
	if r < n && arr[largest].Count < arr[r].Count {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapifyItem(arr, n, largest)
	}
}
