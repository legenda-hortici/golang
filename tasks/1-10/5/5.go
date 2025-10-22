package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
	}
}

func main() {
	var wg sync.WaitGroup

	jobs := make(chan int, 100)

	for j := 0; j < 100; j++ {
		jobs <- j
	}

	close(jobs)

	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	wg.Wait()
}
