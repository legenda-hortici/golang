package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int         // заводим переменную счетчик
	var wg sync.WaitGroup // объявляем WaitGroup
	mu := sync.Mutex{}    // объявляем Mutex

	for i := 1; i <= 1000; i++ { // создаем 1000 горутин
		wg.Add(1) // обновляем очередь горутин
		go func() {
			defer wg.Done() // убираем горутину из очереди (-1)
			mu.Lock()       // блокируем мьютекс
			{
				count++
			}
			mu.Unlock() // разблокировываем мьютекс
		}()
	}

	wg.Wait() // блокируем главный поток
	fmt.Printf("goroutine count: %d\n", count)
}
