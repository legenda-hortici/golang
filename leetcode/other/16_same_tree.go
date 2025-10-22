package other

import "errors"

/*
Если даны корни двух бинарных деревьев p и q, напишите функцию, проверяющую, одинаковы они или нет.
Два бинарных дерева считаются одинаковыми, если они структурно идентичны, а их узлы имеют одинаковое значение.
*/
// https://leetcode.com/problems/same-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	if t.Val == value {
		return errors.New("node already exists")
	}

	if t.Val > value {
		if t.Left == nil {
			t.Left = &TreeNode{Val: value}
			return nil
		}
		return t.Left.Insert(value)
	}

	if t.Val < value {
		if t.Right == nil {
			t.Right = &TreeNode{Val: value}
			return nil
		}
		return t.Right.Insert(value)
	}

	return nil
}
