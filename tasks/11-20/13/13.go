package main

import "fmt"

/*
Напиши функцию Reverse(slice []int), которая разворачивает слайс "на месте" (in-place), без выделения нового массива.
*/

func reverse(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	return nums
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("original slice: %d\n", slice)

	fmt.Printf("reverse slice: %d\n", reverse(slice))
}
