package main

import (
	"sync"
	"time"
)

//===========================================================
//Задача 18
//1. Иногда приходят нули. В чем проблема? Исправь ее
//2. Если функция bank_network_call выполняется 5 секунд, то за сколько выполнится balance()? Как исправить проблему?
//3. Представим, что bank_network_call возвращает ошибку дополнительно. Если хотя бы один вызов завершился с ошибкой, то balance должен вернуть ошибку.
//===========================================================

func bank_network_call(i int) (int, error) {
	time.Sleep(5 * time.Second)
	return i * 2, nil
}

func balance() (int, error) {
	x := make(map[int]int, 1)
	errCh := make(chan error, 1)

	var m sync.Mutex
	var wg sync.WaitGroup

	// call bank
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			b, err := bank_network_call(i)
			if err != nil {
				errCh <- err
				return
			}

			m.Lock()
			x[i] = b
			m.Unlock()
		}(i)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return 0, err
		}
	}

	sumOfMap := 0
	for _, val := range x {
		sumOfMap += val
	}

	// Как-то считается сумма значений в мапе и возвращается
	return sumOfMap, nil
}

func main() {

}
