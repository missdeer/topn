package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	h1 := hash("a")
	if h1 != 97 {
		t.Error(`hash("a") == 97`)
	}
	h2 := hash("abc")
	if h2 != 34 {
		t.Error(`hash("abc") == 34`)
	}
	h3 := hash("https://www.google.com")
	if h3 != 37 {
		t.Error(`hash("https://www.google.com") == 37`)
	}
	h4 := hash("https://minidump.info")
	if h4 != 57 {
		t.Error(`hash("https://minidump.info") == 57`)
	}
	h5 := hash("https://github.com")
	if h5 != 117 {
		t.Error(`hash("https://github.com") == 117`)
	}
}

func TestHash2(t *testing.T) {
	h1 := hash2("a")
	if h1 != 97 {
		t.Error(`hash("a") == 97`, h1)
	}
	h2 := hash2("abc")
	if h2 != 3 {
		t.Error(`hash("abc") == 3`, h2)
	}
	h3 := hash2("https://www.google.com")
	if h3 != 117 {
		t.Error(`hash("https://www.google.com") == 117`, h3)
	}
	h4 := hash2("https://minidump.info")
	if h4 != 122 {
		t.Error(`hash("https://minidump.info") == 122`, h4)
	}
	h5 := hash2("https://github.com")
	if h5 != 104 {
		t.Error(`hash("https://github.com") == 104`, h5)
	}
}
