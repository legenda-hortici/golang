package main

import (
	"fmt"
	"sort"
)

//==========================================
//Задача 1
//Что выведет код?
//==========================================

func main() {
	v := []int{3, 4, 1, 2, 5}
	ap(v)          // 3 4 1 2 5
	sr(v)          // 3 4 1 2 5
	fmt.Println(v) // 1 2 3 4 5
}

func ap(arr []int) {
	// создается новый слайс тк cap переполнен
	arr = append(arr, 10) // 3 4 1 2 5 10
}

func sr(arr []int) {
	sort.Ints(arr) // 1 2 3 4 5
}
