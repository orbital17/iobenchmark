package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solveScanner()
}

func solveScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var n, v, sum int32
	scanner.Scan()
	n = toInt(scanner.Bytes())
	for i := 0; i < int(n); i++ {
		scanner.Scan()
		v = toInt(scanner.Bytes())
		sum += v
	}
	fmt.Fprintf(os.Stdout, "%d\n", sum)
}

func toInt(buf []byte) (n int32) {
	for _, v := range buf {
		n = n*10 + int32(v-'0')
	}
	return
}
