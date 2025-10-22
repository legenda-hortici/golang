package main

import "fmt"

/*
Реализуй функцию Remove(slice []int, i int) []int, которая удаляет элемент с индексом i.
👉 Подумай: стоит ли использовать append(slice[:i], slice[i+1:]...)?
👉 Объясни, почему это может быть неэффективно на больших слайсах.
*/

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := make([]int, 0, 5)

	for i := 0; i < 100; i++ {
		slice = append(slice, i)
	}

	fmt.Printf("original slice: %d\n", slice)

	fmt.Printf("slice after remove: %d\n", remove(slice, 50))
}
