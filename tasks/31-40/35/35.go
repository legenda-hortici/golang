package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		need := target - num
		if idx, found := seen[need]; found {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
