package main

import (
	"fmt"
	"math"
	"sync"
)

func isPowerOfTwo(n int) bool {
	var wg sync.WaitGroup
	wg.Add(1)
	flag := false
	go func() {
		defer wg.Done()
		for i := 0; i <= n; i++ {
			if n == int(math.Pow(2.0, float64(i))) {
				flag = true
			}
		}
	}()
	wg.Wait()
	return flag
}

func main() {
	fmt.Println(isPowerOfTwo(16777215))
}
