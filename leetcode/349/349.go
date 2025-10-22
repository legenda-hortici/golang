package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	result := make([]int, 0)

	var m map[int]bool
	var big []int
	if len(nums1) >= len(nums2) {
		m = make(map[int]bool)
		for i := 0; i < len(nums2); i++ {
			m[nums2[i]] = true
		}
		big = nums1
	} else if len(nums2) >= len(nums1) {
		m = make(map[int]bool)
		for i := 0; i < len(nums1); i++ {
			m[nums1[i]] = true
		}
		big = nums2
	}

	for i := 0; i < len(big); i++ {
		if m[big[i]] {
			result = append(result, big[i])
		}
	}

	clean := make(map[int]bool)
	res := make([]int, 0)
	for i := 0; i < len(result); i++ {

		if !clean[result[i]] {
			res = append(res, result[i])
			clean[result[i]] = true
		}

	}

	return res
}

func main() {
	fmt.Println(intersection([]int{1}, []int{1}))

}
