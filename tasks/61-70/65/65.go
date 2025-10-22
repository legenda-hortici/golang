package main

import (
	"fmt"
	"math/rand"
)

//==========================================
//Задача 1
//1. Написать функцию, которая принимает число N и возвращает слайс размера N с уникальными числами.
//2. Идеи как тестировать функцию?
//==========================================

func create(n int) []int {
	uniq := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		val := rand.Intn(20)
		uniq[val]++
		if uniq[val] == 1 {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	count := 10
	fmt.Println(count, create(count))
}
