package main

import (
	"fmt"
	"sync"
)

//===========================================================
//Задача 9
//1. Что выведется и как исправить?
//2. Что поправить, чтобы сохранить порядок?
//===========================================================

type Msg struct {
	Idx int
	Msg string
}

func main() {
	m := make(chan Msg, 3) // создаем буф канал
	var wg sync.WaitGroup
	cnt := 5
	for i := 0; i < cnt; i++ { // запускаем 5 горутин
		wg.Add(1)
		go func(val int) { // 2 оставшиеся горутины повиснут тк буфер будет забит пока не прочитают
			defer wg.Done()
			m <- Msg{
				Idx: i,
				Msg: fmt.Sprintf("Goroutine: %d\n", val),
			} // пытаемся записать строки в канал
		}(i)
	}

	go func() {
		wg.Wait()
		close(m)
	}()

	//for i := 0; i < cnt; i++ { // 5 раз рапускам горутину на чтение
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		ReceiveFromCh(m, cnt)
	//	}()
	//}
	//wg.Wait()
	results := make([]string, cnt)
	for msg := range m {
		results[msg.Idx] = msg.Msg
	}

	for _, r := range results {
		fmt.Println(r)
	}
}

//func ReceiveFromCh(ch chan Msg, cnt int) {
//	results := make([]string, cnt)
//	for v := range ch {
//		results[v.Idx] = v.Msg
//	}
//	for _, result := range results {
//		fmt.Println(result)
//	}
//	//fmt.Println(<-ch) // читаем весь канал
//}
