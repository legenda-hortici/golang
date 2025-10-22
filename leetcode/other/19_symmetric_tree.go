package other

// Задав корень бинарного дерева, проверьте, является ли оно зеркальным отражением самого себя (т.е. симметричным относительно центра).
// https://leetcode.com/problems/symmetric-tree/description/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}
