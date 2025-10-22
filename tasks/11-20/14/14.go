package main

import "fmt"

/*
Напиши код, который показывает, как append влияет на len и cap.
Пример: создай слайс длиной 3, добавь несколько элементов и выведи capacity после каждого добавления.
*/

func main() {
	slice := make([]int, 3)

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("after append: len = %d, cap = %d\n", len(slice), cap(slice))
	}
}

/*
	after append: len = 4, cap = 6
	after append: len = 5, cap = 6
	after append: len = 6, cap = 6
	after append: len = 7, cap = 12
	after append: len = 8, cap = 12
	after append: len = 9, cap = 12
	after append: len = 10, cap = 12
	after append: len = 11, cap = 12
	after append: len = 12, cap = 12
	after append: len = 13, cap = 24
*/
