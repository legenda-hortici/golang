package main

import "fmt"

func main() {
	numbers := make([]*int, 0, 5)
	var number int
	for range 3 {
		number++
		k := 2
		numbers = append(numbers, &number, &k)
	}

	double(numbers)

	for _, num := range numbers {
		fmt.Printf("%d ", *num) // 6 4 6 4 6 4
	}
}

func double(numbers []*int) {
	*numbers[0] *= 2
	for i := 1; i < len(numbers); i++ {
		if i%2 != 0 {
			*numbers[i] = *numbers[i] * 2
		}
	}
}
