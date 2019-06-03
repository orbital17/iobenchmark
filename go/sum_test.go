package main

import (
	"bufio"
	"bytes"
	"fmt"
	"runtime"
	"testing"
)

func analyzeMemory(f func()) {
	var old, new runtime.MemStats
	runtime.ReadMemStats(&old)
	f()
	runtime.ReadMemStats(&new)
	fmt.Printf("used memory: %v Mb\n", bToMb(new.TotalAlloc-old.TotalAlloc))
}

func bToMb(b uint64) float64 {
	return float64(b/1024) / 1024
}

var sizes = map[string]struct {
	Filename string
	Result   string
}{
	"s":  {"input", "314\n"},
	"m":  {"inputmedium", "4717353\n"},
	"ml": {"inputml", "37750341\n"},
	"l":  {"inputlarge", "301998100\n"},
}

func Test_makeIO(t *testing.T) {
	s := sizes["ml"]
	var buffer bytes.Buffer
	reader := fileReader(s.Filename)
	writer := bufio.NewWriter(&buffer)
	analyzeMemory(func() { makeIOAndSolve(reader, writer) })
	result := buffer.String()
	if result != s.Result {
		t.Fatalf("got %v", result)
	}
}

func Benchmark_makeIO(b *testing.B) {
	s := sizes["m"]
	for i := 0; i < b.N; i++ {
		reader := fileReader(s.Filename)
		var buffer bytes.Buffer
		writer := bufio.NewWriter(&buffer)
		makeIO(reader, writer)
		b.StopTimer()
		result := buffer.String()
		if result != s.Result {
			b.Fatalf("got %v", result)
		}
		b.StartTimer()
	}
}
