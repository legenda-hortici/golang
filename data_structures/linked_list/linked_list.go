package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (l *LinkedList) AppendToBack(value int) {
	newNode := &ListNode{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (l *LinkedList) DeleteFromBack() {
	if l.head == nil {
		return
	}

	current := l.head

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
}

func (l *LinkedList) GetLast() {
	if l.head == nil {
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	fmt.Println(current.value, nil)
}

func (l *LinkedList) GetByValue(value int) {
	if l.head == nil {
		return
	}

	current := l.head
	for current.next != nil {
		if current.value == value {
			fmt.Printf("%d\n", value)
			return
		}
		current = current.next
	}
}

func (l *LinkedList) Print() {
	current := l.head

	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println()
}

func main() {
	var node LinkedList

	node.AppendToBack(1)
	node.AppendToBack(2)
	node.AppendToBack(3)

	node.Print()
	fmt.Print("last node value: ")
	node.GetLast()

	node.DeleteFromBack()
	node.Print()

	fmt.Print("last node value: ")
	node.GetLast()

	fmt.Print("node with value: ")
	node.GetByValue(2)

}
