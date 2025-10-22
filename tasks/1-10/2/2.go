package main

import (
	"fmt"
	"time"
)

func main() {
	helloChan := make(chan string) // создали канал

	go func() {
		defer close(helloChan)     // закрываем через defer
		helloChan <- "Hello World" // записали переменную
	}()

	time.Sleep(1 * time.Second) // ждем секунду, чтобы успеть записать

	go func() {
		fmt.Println(<-helloChan) // выводим в консоль
	}()

	time.Sleep(1 * time.Second) // ждем, чтобы успеть прочитать
}
