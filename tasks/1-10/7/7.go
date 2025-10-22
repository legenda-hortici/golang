package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go generate(chan1, wg)

	go multiply(chan1, chan2, wg)

	go printResult(chan2, wg)

	wg.Wait()
}

func generate(out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	for i := 0; i < 5; i++ {
		num := rand.Intn(10)
		fmt.Printf("Generated: %d\n", num)
		out <- num
	}
}

func multiply(out, chan2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(chan2)

	for v := range out {
		fmt.Printf("Multiplied: %d\n", v*2)
		chan2 <- v * 2
	}
}

func printResult(chan2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range chan2 {
		fmt.Printf("Printed: %d\n", v)
	}
}
