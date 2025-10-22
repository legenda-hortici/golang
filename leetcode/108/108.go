package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{nums[mid], nil, nil}

	root.left = sortedArrayToBST(nums[:mid])
	root.right = sortedArrayToBST(nums[mid+1:])

	return root
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	root := sortedArrayToBST(array)
	fmt.Println(root)
}
