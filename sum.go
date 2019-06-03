package main

import (
	"bufio"
	"fmt"
)

func simpleArraySum(ar []int32) int32 {
	sum := int32(0)
	for _, v := range ar {
		sum += v
	}
	return sum
}

func makeIO(reader *bufio.Reader, writer *bufio.Writer) {
	var arCount int32
	fmt.Fscan(reader, &arCount)

	var ar []int32 = make([]int32, arCount)
	for i := 0; i < int(arCount); i++ {
		fmt.Fscan(reader, &ar[i])
	}

	result := simpleArraySum(ar)

	fmt.Fprintf(writer, "%d\n", result)
	writer.Flush()
}

func makeIOAndSolve(reader *bufio.Reader, writer *bufio.Writer) {
	var n, v, sum int32
	fmt.Fscan(reader, &n)
	for i := 0; i < int(n); i++ {
		fmt.Fscan(reader, &v)
		sum += v
	}
	fmt.Fprintf(writer, "%d\n", sum)
	writer.Flush()
}
