package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	ctx := context.Background()
	go func() {
		for i := 1; i <= 100; i++ {
			in <- i
		}
		close(in)
	}()

	t := time.Now()
	processParallel(in, out, 5, ctx)

	for v := range out {
		fmt.Println(v)
	}

	fmt.Println(time.Since(t).String())
}

func processData(out chan<- int, in <-chan int, bg context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range in {
		ctx, cancel := context.WithTimeout(bg, time.Second*5)
		done := make(chan int, 1)

		go func() {
			defer cancel()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			done <- val * 2
		}()

		select {
		case <-ctx.Done():
			fmt.Println("done: timeout ", val)
		case res := <-done:
			out <- res
		}

		cancel()
	}
}

// операция не должна выполняться более 5 секунд
func processParallel(in <-chan int, out chan<- int, numWorkers int, ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go processData(out, in, ctx, wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}
