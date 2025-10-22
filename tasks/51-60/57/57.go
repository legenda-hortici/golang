package main

import "fmt"

func main() {
	//==========================================
	//Задача 2
	//1. Что выведет код?
	//==========================================

	var foo []int // 0
	var bar []int // 0

	foo = append(foo, 1) // 1
	foo = append(foo, 2) // 1 2
	foo = append(foo, 3) // 1 2 3
	bar = append(foo, 4) // bar = 1 2 3 4 foo = 1 2 3
	foo = append(foo, 5) // bar = 1 2 3 5 <- boo = 1 2 3 5

	fmt.Println(foo, bar)
	// foo = 1 2 3 5 bar = 1 2 3 5
}
