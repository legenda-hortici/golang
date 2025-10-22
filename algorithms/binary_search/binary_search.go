package main

import "fmt"

func binarySearch(in []int, target int) (int, bool) {
	if len(in) == 0 {
		return 0, false
	}

	first, last := 0, len(in)-1
	for first <= last {
		mid := ((last - first) / 2) + first
		if in[mid] == target {
			return mid, true
		} else if in[mid] < target { // нужно искать в "левой" части слайса
			last = mid - 1
		} else if in[mid] > target { // нужно искать в "правой" части слайса
			first = mid + 1
		}
	}
	return 0, false
}

func main() {
	fmt.Println(binarySearch([]int{0, 3, 7, 2, 4}, 2))
}
