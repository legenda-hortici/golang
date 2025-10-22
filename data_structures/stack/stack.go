package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return 0
	}

	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Get() []int {
	return s.data
}

func main() {
	var stack Stack

	fmt.Println("not initialized stack", stack.Get())
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("pushed stack:", stack.Get())

	stack.Pop()
	fmt.Println("poped from stack:", stack.Get())

	fmt.Println("top of stack:", stack.Top())
}
