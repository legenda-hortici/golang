package main

import "fmt"

func main() {
	var numbers []*int
	for _, num := range []int{1, 2, 3, 4, 5} {
		numbers = append(numbers, &num)
	}

	for _, num := range numbers {
		fmt.Println(num)
	}
}
