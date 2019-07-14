package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	outputFileSize int
	outputFilePath string
)

func main() {
	flag.IntVar(&outputFileSize, "size", 1000*1000*1000, "output file size")
	flag.StringVar(&outputFilePath, "output", "output.txt", "output file path")
	flag.Parse()

	if outputFileSize < 10*1000*1000 {
		log.Fatal("output file size is expected to be larger")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	dedicatesCount := 65535 * int(math.Log10(float64(outputFileSize))-5)
	dedicates := make([]string, dedicatesCount)
	for i := 0; i < dedicatesCount; i++ {
		dedicates[i] = strings.Repeat(fmt.Sprintf("%d", rand.Intn(65535)), rand.Intn(dedicatesCount/65535)+50) + "\n"
	}

	log.Println("creating output file", outputFilePath)
	f, e := os.Create(outputFilePath)
	if e != nil {
		log.Fatal(e)
		os.Exit(2)
	}
	defer f.Close()

	log.Println("generating data, expected to be not less than", outputFileSize, "bytes")
	c, s := 0, 0
	m := make(map[int]int, dedicatesCount)
	for ; s < outputFileSize; c++ {
		index := rand.Intn(dedicatesCount)
		s += len(dedicates[index])
		f.WriteString(dedicates[index])
		m[index]++
	}
	log.Println("done!", s, "bytes", c, "lines has been written.")
}
