package main

import (
	"fmt"
	"time"
)

// длительность выполнения

const numRequests = 10000

var count int

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса
	count++
}

func main() {
	/*
		В Go time.Sleep(d) гарантирует не меньше d, но может быть больше.
		Планировщик ОС «будит» горутину после сна, но не строго через 1 ms, а как получится (обычно 1–2 ms).
		Это называется scheduling latency.
	*/
	start := time.Now()
	for i := 0; i < numRequests; i++ {
		networkRequest()
	}

	fmt.Println(count, time.Since(start)) // 10k ms?
}
