package main

import "fmt"

/*
Реализуй функцию, которая копирует один слайс в другой.
Используй copy(dst, src) и объясни, что произойдет, если размер dst меньше, чем src.
*/

func copySlice(slice []int, newSlice []int) []int {
	copy(slice, newSlice)
	return slice
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5} // len = 5, cap = 5

	slice2 := make([]int, 2)

	newSlice1 := copySlice(slice1, []int{100, 200, 300, 400, 500}) // 100 200 300 400 500
	// так как длина и вместимость совпадает у обоих слайсов

	newSlice2 := copySlice(slice2, []int{100, 200, 300, 400, 500}) // 100 200
	// так как мы пытаемся перекопировать слайс большей длины с меньшей длиной

	fmt.Println(newSlice1)

	fmt.Println(newSlice2)
}
