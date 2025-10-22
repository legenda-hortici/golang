package main

import "fmt"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	check := make(map[int]int)

	for _, v := range nums {
		check[v]++
	}

	for _, v := range nums {
		if check[v] == 1 {
			return v
		}
	}

	return 0
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
