package main

import (
	"fmt"
	"sync"
)

func main() {
	million := make([]int, 0)
	var wg sync.WaitGroup

	for i := 0; i < 1000000; i++ {
		million = append(million, i)
	}

	parts := make(map[int][]int)
	key := 0
	for i := 0; i < len(million); i++ {
		if million[i]%100000 == 0 {
			wg.Add(1)
			parts[key] = million[i : i+100000]

			go findMax(parts[key], &wg)

			key += 1
		}
	}
	wg.Wait()
}

func findMax(part []int, wg *sync.WaitGroup) {
	defer wg.Done()
	maxVal := 0

	for j := 0; j < len(part); j++ {
		if part[j] > maxVal {
			maxVal = part[j]
		}
	}
	fmt.Println(maxVal)
}
