//===========================================================
//Задача 19
//Что выведет код и почему?
//===========================================================

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// есть data race потому что ограничиваем кол-во горутин на 1 ядре: 1 ядро = 1 горутина
	// в данном случае выполняется main как главная горутина, но анонимная горутина попытается захватить значение
	// ch и возникает data race
	runtime.GOMAXPROCS(1) // 1 ядро
	ch := 0

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch = 1
	}()
	wg.Wait()
	for ch == 0 { // busy loop
		//runtime.Gosched()
	}
	fmt.Println("finish")
}
