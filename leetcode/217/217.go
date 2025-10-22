package main

import "fmt"

func containsDuplicate(nums []int) bool {
	check := make(map[int]int)
	for _, n := range nums {
		check[n]++
	}

	for _, v := range check {
		if v > 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
