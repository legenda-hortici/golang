package main

import "fmt"

//===========================================================
//Задача 4
//1. Что выведет код?
//===========================================================

func main() {
	var m map[string]int // nil map
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++ // panic!
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
