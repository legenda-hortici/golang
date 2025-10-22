package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan int)

	// go func() {
	// 	channel <- 42
	// }()

	// close(channel)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		select {
		case _, ok := <-channel:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("received")
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
		}
	}()
	// time.Sleep(3 * time.Second)

	wg.Wait()
}
