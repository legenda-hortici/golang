package main

import (
	"fmt"
	"sync"
)

func FanOut(bigChan chan int, num int) []chan int {
	res := make([]chan int, num)

	for k := 0; k < len(res); k++ {
		res[k] = make(chan int)
	}

	go func() {
		defer func() {
			for _, ch := range res {
				close(ch)
			}
		}()
		i := 0
		for val := range bigChan {
			res[i] <- val // 1: 1 4, 2: 2 5, 3: 3 6
			i++           // 1 2 3
			if i == len(res) {
				i = 0
			}
		}
	}()

	return res
}

func main() {
	bigChan := make(chan int)

	go func() {
		defer close(bigChan)
		for i := 0; i < 10; i++ {
			bigChan <- i
		}
	}()

	res := FanOut(bigChan, 3)

	var wg sync.WaitGroup
	wg.Add(len(res))

	for i, ch := range res {
		go func(ch chan int, i int) {
			defer wg.Done()
			for val := range ch {
				fmt.Printf("channel %d: %d\n", i+1, val)
			}
		}(ch, i)
	}

	wg.Wait()
}
