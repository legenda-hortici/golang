package main

import (
	"fmt"
	"sync"
)

//===========================================================
//Задача 6
//1. Что выведет код и как исправить?
//===========================================================

var globalMap = map[string][]int{
	"test":  make([]int, 0),
	"test2": make([]int, 0),
	"test3": make([]int, 0),
}
var a = 0

func main() {
	/*
		данные в мапе будут разные, потому что переменная "а" общая для всех
		чтобы это исправить можно передавать ее как параметр в горутины или сделать
		локальные переменные в самих горутинах. а также добавить мьютексы на lock-unlock
		при добавлении элемента в мапу
	*/
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	
	wg.Add(3)
	go func(a int) {
		wg.Done()
		a = 10
		mu.Lock()
		globalMap["test"] = append(globalMap["test"], a)
		mu.Unlock()
	}(a)
	go func(a int) {
		wg.Done()
		a = 11
		mu.Lock()
		globalMap["test2"] = append(globalMap["test2"], a)
		mu.Unlock()
	}(a)
	go func(a int) {
		wg.Done()
		a = 12
		mu.Lock()
		globalMap["test3"] = append(globalMap["test3"], a)
		mu.Unlock()
	}(a)
	wg.Wait()
	fmt.Printf("%v\n", globalMap)
	fmt.Printf("%d", a)
}
