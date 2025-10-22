package main

import (
	"fmt"
	"math/rand"
)

/*
Напиши функцию, которая принимает массив []int и возвращает новый слайс, где все элементы умножены на 2.
👉 Важно: не изменяй оригинальный слайс.
*/

func multy(nums []int) []int {
	newSlice := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newSlice[i] = nums[i] * 2
	}

	return newSlice
}

func main() {
	nums := make([]int, 0)

	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(10))
	}

	fmt.Printf("original slice: %d\n", nums)

	newSlice := multy(nums)

	fmt.Printf("slice after multy: %d\n", newSlice)

	fmt.Println(nums)
}
