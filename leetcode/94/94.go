package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)              // обходим левое поддерево
		res = append(res, node.Val) // записываем значение посещенного узла
		dfs(node.Right)             // обходим правое поддерево
	}

	dfs(root)
	return res
}

func main() {

}
