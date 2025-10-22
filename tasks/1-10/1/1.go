package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Printf("%d ", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Printf("%d ", i)
			}
		}
	}()

	wg.Wait()
}
