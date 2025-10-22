package main

import "fmt"

func majorityElement(nums []int) int {
	check := make(map[int]int)

	for _, n := range nums {
		check[n]++
	}

	fmt.Println(check)

	for k, val := range check {
		fmt.Println(k, val)
		if val > len(nums)/2 {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(majorityElement([]int{2, 3, 2}))
}
