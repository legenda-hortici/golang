package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Push(value int) {
	q.data = append(q.data, value)
}

func (q *Queue) Pop() int {
	if len(q.data) == 0 {
		return 0
	}

	first := q.data[0]
	q.data = q.data[first:]
	return first
}

func (q *Queue) Print() {
	fmt.Println(q.data)
}

func (q *Queue) First() int {
	return q.data[0]
}

func main() {
	var queue Queue

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Print()

	queue.Pop()
	queue.Print()

	queue.Push(5)
	queue.Print()

	fmt.Println(queue.First())

}
