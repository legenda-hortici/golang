package main

import "fmt"

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- f(j)
	}
}

func main() {
	numJobs := 100
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	multiplier := func(i int) int {
		return i * 10
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for w := 1; w <= 3; w++ {
		go worker(multiplier, jobs, results)
	}

	for r := 1; r <= numJobs; r++ {
		fmt.Println(<-results)
	}
}
