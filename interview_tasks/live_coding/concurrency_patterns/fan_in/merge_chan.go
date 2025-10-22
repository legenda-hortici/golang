package main

import (
	"fmt"
	"sync"
)

/*
	Требуется объединить n количество каналов в один результирующий канал и вывести его
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 11; i <= 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 21; i <= 30; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	mergedChan := merge(ch1, ch2, ch3)

	for i := range mergedChan {
		fmt.Println(i)
	}
}

func merge(chans ...chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range chans {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for i := range ch {
				out <- i
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	//wg.Wait()
	//close(out)

	return out
}
