package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

func main() {
	fmt.Println(missingNumber([]int{0, 1}))
}
