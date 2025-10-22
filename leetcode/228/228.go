package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	start := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if start == nums[i-1] {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i]
		}
	}

	if start == nums[len(nums)-1] {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, nums[len(nums)-1]))
	}

	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
