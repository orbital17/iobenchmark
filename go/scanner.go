package main

import (
	"bufio"
	"fmt"
	"io"
)

func solveScanner(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	var n, v, sum int32
	scanner.Scan()
	n = toInt(scanner.Bytes())
	for i := 0; i < int(n); i++ {
		scanner.Scan()
		v = toInt(scanner.Bytes())
		sum += v
	}
	fmt.Fprintf(writer, "%d\n", sum)
}

func toInt(buf []byte) (n int32) {
	for _, v := range buf {
		n = n*10 + int32(v-'0')
	}
	return
}
