package main


func hash(s string) (h int) {
	for i := 0; i < len(s); i++ {
		h += h<<5 - h + int(s[i])
	}
	h ^= (h >> 16)
	h &= (N - 1)
	return h
}
