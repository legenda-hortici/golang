package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))

}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
		if i >= k {
			delete(m, nums[i-k])
			m[nums[i]] = true
		}
	}
	return false
}
