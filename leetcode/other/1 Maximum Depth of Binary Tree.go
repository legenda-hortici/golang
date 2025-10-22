package other

/*
Задав корень бинарного дерева, верните его максимальную глубину.

Максимальная глубина бинарного дерева - это количество узлов на самом длинном пути от корневого узла до самого дальнего листового узла.
*/

// https://leetcode.com/problems/maximum-depth-of-binary-tree/submissions/1591469572/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
