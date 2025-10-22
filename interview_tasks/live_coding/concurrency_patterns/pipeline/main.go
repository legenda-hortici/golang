package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	writeCh := make(chan int)
	squareCh := make(chan int)
	decreaseCh := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			writeCh <- i
			fmt.Printf("goriutine 1: %d\n", i)
		}
		close(writeCh)
	}()

	go func() {
		for val := range writeCh {
			squareCh <- val * val
			fmt.Printf("goriutine 2: %d\n", val*val)
		}
		close(squareCh)
	}()

	go func() {
		for val := range squareCh {
			decreaseCh <- val - 10
			fmt.Printf("goriutine 3: %d\n", val-10)
		}
		close(decreaseCh)
	}()

	for val := range decreaseCh {
		fmt.Println(val)
	}

	fmt.Printf("%s lapsed\n", time.Since(start))
}
