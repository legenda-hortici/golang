package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func reverseBits(n int) int {
	bits := fmt.Sprintf("%032b", n)

	var buf bytes.Buffer
	for i := len(bits) - 1; i >= 0; i-- {
		buf.WriteByte(bits[i])
	}

	newNum, _ := strconv.ParseInt(buf.String(), 2, 32)
	return int(newNum)
}

func main() {
	fmt.Println(reverseBits(43261596))
}
