package main

import "fmt"

func main() {

	nums := make([]int, 1, 3)
	fmt.Println(nums) // 0 тк len = 1

	appendSlice(nums, 1) // передача по копии
	fmt.Println(nums)    // 0 тк len = 1

	copySlice(nums, []int{2, 3, 4, 5, 6, 7, 8}) // передача по копии
	fmt.Println(nums)                           // 2 тк исходный nums len = 1 поэтому из копии достанется только первый элемент

	mutateSlice(nums, 1, 4)
	fmt.Println(nums) // паника index out of range
}

func appendSlice(slice []int, num int) {
	slice = append(slice, num) // 0 1
}

func copySlice(slice, newSlice []int) {
	// len(slice) = 1; len(newSlice) = 2 | 2 -> 1 (!!!)
	copy(slice, newSlice) // newSlice len = 2, slice len = 2
	fmt.Println("old:", slice, "\nnew:", newSlice)
}

func mutateSlice(slice []int, idx, val int) {
	slice[idx] = val
}
