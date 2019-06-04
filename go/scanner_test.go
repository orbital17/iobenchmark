package main

import (
	"bufio"
	"bytes"
	"testing"
)

var sizes = map[string]struct {
	Filename string
	Result   string
}{
	"s":  {"../input", "314\n"},
	"m":  {"../inputmedium", "4717353\n"},
	"ml": {"../inputml", "37750341\n"},
	"l":  {"../inputlarge", "301998100\n"},
}

func Test_solveScanner(t *testing.T) {
	s := sizes["l"]
	var buffer bytes.Buffer
	reader := fileReader(s.Filename)
	writer := bufio.NewWriter(&buffer)
	analyzeMemory(func() { solveScanner(reader, writer) })
	writer.Flush()
	result := buffer.String()
	if result != s.Result {
		t.Fatalf("got %v", result)
	}
}

func Benchmark_solveScanner(b *testing.B) {
	s := sizes["l"]
	for i := 0; i < b.N; i++ {
		reader := fileReader(s.Filename)
		var buffer bytes.Buffer
		writer := bufio.NewWriter(&buffer)
		solveScanner(reader, writer)
		writer.Flush()
		b.StopTimer()
		result := buffer.String()
		if result != s.Result {
			b.Fatalf("got %v", result)
		}
		b.StartTimer()
	}
}
