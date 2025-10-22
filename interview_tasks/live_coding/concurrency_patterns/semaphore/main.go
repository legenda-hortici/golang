package main

import (
	"fmt"
)

type Sema chan struct{}

func NewSema(n int) Sema {
	return make(chan struct{}, n)
}

func (s Sema) Add(count int) {
	for i := 0; i < count; i++ {
		s <- struct{}{}
	}
}

func (s Sema) Wait(count int) {
	for i := 0; i < count; i++ {
		<-s
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(nums)

	semaphore := NewSema(n)

	for i := 0; i < n; i++ {
		go func() {
			fmt.Printf("goroutine %d: %d\n", i, nums[i])
			semaphore.Add(1)
		}()
	}

	semaphore.Wait(n)
}
