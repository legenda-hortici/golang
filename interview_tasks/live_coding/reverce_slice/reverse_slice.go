package main

import "fmt"

/*
	Реализовать функцию reverse, разворачивающую срез целых чисел без использования временного среза
*/

func reverse(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
		// 0, 2 = 2, 0 for [0 1 2] => [2 1 0]

		// 0, 3 = 3, 0 for [0 1 2 3] => [3 1 2 0]
		// 1, 2 = 2, 1 for [0 1 2 3] => [3 2 1 0]
	}
	return slice
}

func main() {
	nums := []int{0, 1, 2, 3}
	fmt.Println("before", nums)
	reverse(nums)
	fmt.Println("after", nums)
}
