package main

import (
	"bufio"
	"os"
)

func readerWriterLocal() (*bufio.Reader, *bufio.Writer) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	return reader, writer
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func fileReader(filename string) *bufio.Reader {
	f, _ := os.Open(filename)
	return bufio.NewReader(f)
}
