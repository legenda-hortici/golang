package main

import (
	"fmt"
	"time"
)

// что будет, сколько горутин

// panic e
// goroutines: 2

func main() {
	// recover() может поймать панику только в той же горутине, где стоит defer.
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("was get panic %v\n", err)
		}
	}()
	go ff() // panic: e не продолжаем дальше

	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		fmt.Printf("work continue")
	}
}

func ff() {
	// По умолчанию Go: если хоть одна горутина упала с паникой и не восстановилась. то крашится вся программа.
	panic("e")
}
