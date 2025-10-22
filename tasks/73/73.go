package main

import "fmt"

// ===========================================================
// Задача 3
// Исправить функцию, чтобы она работала. Сигнатуру менять нельзя
// ===========================================================

func printNumber(ptrToNumber interface{}) {
	if ptrToNumber.(*int) != nil {
		fmt.Println(*ptrToNumber.(*int))
	} else {
		fmt.Println("nil")
	}
}

func main() {
	v := 10         // 10
	printNumber(&v) // 0x000... -> 10
	var pv *int     // nil
	printNumber(pv) // nil -> nil
	pv = &v         // 10
	printNumber(pv) // 0x000... -> 10
}
