package main

import "fmt"

//===========================================================
//Задача 8
//1. Расскажи подробно что происходит
//===========================================================

// # Вариант 1
// -----------
func main() {
	a := []int{1, 2}  // len = 2 cap = 2
	a = append(a, 3)  // len = 3 cap = 4 // 1 2 3
	b := append(a, 4) // len = 4 cap = 4 // 1 2 3 4 -> 1 2 3 5
	c := append(a, 5) // len = 5 cap = 4 // 1 2 3 5

	fmt.Println(b) // 1 2 3 5
	fmt.Println(c) // 1 2 3 5
}

//# Вариант 2
//-----------
//
//func main() {
//	a := []int{1,2}
//	a = append(a, 3)
//	a = append(a, 7)
//	b := append(a, 4)
//	c := append(a, 5)
//
//	fmt.Println(b)
//	fmt.Println(c)
//}
