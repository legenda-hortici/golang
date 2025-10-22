package main

import "fmt"

// ===========================================================
// Задача 9
// Что выведет код и почему?
// ===========================================================
func main() {
	foo := make([]int, 0, 4)   // [] // len = 0 cap = 4
	foo = append(foo, 1, 2, 3) // 1 2 3 // len = 3 cap = 4
	bar := append(foo, 4)      // foo = 1 2 3 4 -> 1 2 3 5 len = 4 cap = 4
	baz := append(foo, 5)      // foo = 1 2 3 5

	fmt.Println(bar) // 1 2 3 5
	fmt.Println(baz) // 1 2 3 5
}
