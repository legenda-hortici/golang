package main

import "fmt"

func main() {
	c := make(chan struct{}, 1)
	// если канал имеет буфер, то он не вернет при чтении zero-value
	// Чтение из открытого пустого канала -> блокировка (канал пуст)
	// Чтение из закрытого канала -> всегда готово, возвращает zero-value + ok=false.
	for done := false; !done; {
		// На 3 итерации, когда несколько кейсов готовы одновременно (default и c <- struct{}{}), Go выбирает случайный.
		select {
		default:
			fmt.Println(1)
			done = true
		case <-c: // чтение из канала (сработает, если в канале есть элемент)
			fmt.Print(2)
			c = nil
		case c <- struct{}{}: // запись в канал (сработает, если в канале есть место)
			fmt.Print(3)
		}
	}
	// 321
}
