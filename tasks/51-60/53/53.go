package main

import (
	"fmt"
	"sync"
	"time"
)

//===========================================================
//Задача 17
//1. Что выведется? Исправь проблему
//===========================================================

func main() {
	x := sync.Map{}
	go func() {
		x.Store(1, 1)
	}()
	go func() {
		x.Store(1, 7)
	}()
	go func() {
		x.Store(1, 10)
	}()
	time.Sleep(100 * time.Millisecond)
	val, ok := x.Load(1)
	if !ok {
		fmt.Println("val is nil")
	}
	fmt.Println("x[1] =", val)
}
