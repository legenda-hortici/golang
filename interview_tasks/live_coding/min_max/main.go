/*
	Реализовать нахождение min и max
*/

package main

import "fmt"

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return b
}

func maxNum(a, b int) int {
	if a < b {
		return b
	}
	return b
}

func main() {
	a, b := 5, 10

	fmt.Println(minNum(a, b), maxNum(a, b))
}
