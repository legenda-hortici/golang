package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.left == nil && root.right == nil {
		if root.val == targetSum {
			return true
		}
	}

	targetSum -= root.val

	return hasPathSum(root.left, targetSum) || hasPathSum(root.right, targetSum)
}

func main() {

}
