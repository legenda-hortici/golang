package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range numbers {
		need := target - num
		if idx, found := seen[need]; found {
			return []int{idx, i + 1}
		}
		seen[num] = i + 1
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
