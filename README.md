# topn
pick up Top 100 items

# Build

```
go get github.com/missdeer/topn
```

# Usage

* use script in directory `gen` (`go get github.com/missdeer/topn/gen`) to generate test data.
* pick up top 100 items from specified text file, eg:
  
```
topn -input testdata/1.txt
```

# Test

```
>go test -v
=== RUN   TestCount
--- PASS: TestCount (0.00s)
=== RUN   TestHash
--- PASS: TestHash (0.00s)
=== RUN   TestHash2
--- PASS: TestHash2 (0.00s)
=== RUN   TestHeapSort
--- PASS: TestHeapSort (0.00s)
=== RUN   TestTopNHeapSort
--- PASS: TestTopNHeapSort (0.00s)
=== RUN   TestSplit
--- PASS: TestSplit (3.13s)
=== RUN   TestSplit2
--- PASS: TestSplit2 (2.99s)
PASS
ok  	github.com/missdeer/topn	6.123s
```

# Notes
Make sure your disk has enough free space for intermediate files during program runs. Otherwise you may specify another directory path to store the temporay files like below:

```
topn -input testdata/1.txt -tempDir /tmp
```
