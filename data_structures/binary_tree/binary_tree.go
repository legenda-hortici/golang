package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("nil tree")
	}

	if t.val == value {
		return errors.New("value already inserted")
	}
	// если вставляемое значение меньше сущ., то проверяется левый узел
	if t.val > value {
		// если левый узел == nil, то инициализируем узел этим значением
		if t.left == nil {
			t.left = &TreeNode{val: value}
			// возвращается пустая ошибка
			return nil
		}
		// если узел лист есть, то вставляем значение рекурсивно
		_ = t.left.Insert(value)
	}
	// если вставляемое значение больше сущ., то проверяется правый узел
	if t.val < value {
		// если правый узел == nil, то инициализируем узел этим значением
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		// если узел лист есть, то вставляем значение рекурсивно
		_ = t.right.Insert(value)
	}
	return nil
}

func (t *TreeNode) FindMin() (int, error) {
	if t == nil {
		return 0, errors.New("nil tree")
	}
	// если левый потомок == nil, то берем значение листа
	if t.left == nil {
		return t.val, nil
	}
	// иначе рекурсивно ищем дальше
	return t.left.FindMin()
}

func (t *TreeNode) FindMax() (int, error) {
	if t == nil {
		return 0, errors.New("nil tree")
	}
	// если правый потомок == nil, то берем значение листа
	if t.right == nil {
		return t.val, nil
	}
	// иначе рекурсивно ищем дальше
	return t.right.FindMax()
}

func (t *TreeNode) Find(value int) *TreeNode {
	if t == nil {
		return nil
	}

	if t.val == value {
		return t
	}

	if t.val > value {
		return t.left.Find(value)
	} else {
		return t.right.Find(value)
	}
}

func (t *TreeNode) Delete(value int) (*TreeNode, error) {
	if t == nil {
		return nil, errors.New("nil tree")
	}

	if value < t.val {
		t.left, _ = t.left.Delete(value)
	} else if value > t.val {
		t.right, _ = t.right.Delete(value)
	} else {
		if t.left == nil && t.right == nil {
			return nil, nil // нет потомков
		} else if t.left == nil {
			return t.right, nil // только правый потомок
		} else if t.right == nil {
			return t.left, nil // только левый потомок
		} else {
			// 2 потомка
			minVal, _ := t.right.FindMin()
			t.val = minVal
			t.right, _ = t.right.Delete(value)
		}
	}

	return t, nil
}

func (t *TreeNode) Print() {
	if t == nil {
		return
	}

	t.left.Print()
	fmt.Printf("node value: %d\n", t.val)
	t.right.Print()
}

func main() {
	tree := &TreeNode{
		val:   1,
		left:  nil,
		right: nil,
	}

	_ = tree.Insert(2)
	_ = tree.Insert(3)
	_ = tree.Insert(1)

	tree.Print()

	if maxVal, err := tree.FindMax(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Max value: %d\n", maxVal)
	}

	if minVal, err := tree.FindMin(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Min value: %d\n", minVal)
	}
	node := tree.Find(3)
	fmt.Printf("Found node: %v\n", *node)

	node, _ = tree.Delete(3)

	tree.Print()
}
