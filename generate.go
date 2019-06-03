package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
)

func getWriter() *bufio.Writer {
	file, err := os.Create("inputml")
	checkError(err)
	return bufio.NewWriter(file)
}

func generateInput() {
	n := int(int32(2 << 22))
	// n := int(int32(2 << 5))
	w := getWriter()
	_, _ = w.WriteString(strconv.Itoa(n) + "\n")
	for i := 0; i < n; i++ {
		_, _ = w.WriteString(strconv.Itoa(rand.Intn(10)) + " ")
	}
	w.Flush()
}
