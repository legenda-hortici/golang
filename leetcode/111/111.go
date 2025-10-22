package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1
	}

	if root.left == nil {
		return 1 + minDepth(root.right)
	}
	if root.right == nil {
		return 1 + minDepth(root.left)
	}

	leftDepth := minDepth(root.left)
	rightDepth := minDepth(root.right)

	if leftDepth < rightDepth {
		return 1 + leftDepth
	}

	return 1 + rightDepth
}

func main() {

}
