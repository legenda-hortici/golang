package main

import "fmt"

func mySqrt(x int) int {
	for i := 1; i <= x; i++ {
		if i*i >= x {
			return i
		} else if i*i >= x && (i-1)*(i-1) < x {
			return i - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(mySqrt(8))
}
