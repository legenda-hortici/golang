package main

import "math"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := height(root.left)
	right := height(root.right) // 20 -> [15 7] (+1); 15 -> [nil nil] (1); 7 -> [nil nil] (1)

	if math.Abs(float64(left-right)) > 1 {
		return false
	}
	return isBalanced(root.left) && isBalanced(root.right)
}

func height(root *TreeNode) int {
	// 20
	if root == nil { // false
		return 0
	}

	return 1 + max(height(root.right), height(root.left)) // + 1 | right: + 1; left: + 1 | right: nil; left: nil: right: nil; left: nil
}

func main() {

}
