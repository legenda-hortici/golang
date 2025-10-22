package main

import "fmt"

func sumOfSquare(c, quit chan int) {
	y := 1 // начальная точка

	for {
		select {
		case c <- (y * y): // записываем результат в канал
			y++ // увеличиваем начальную точку
		case <-quit: // если есть что читать из канала quit, то выходим
			return
		}
	}
}

func main() {
	channel := make(chan int)
	quit := make(chan int)
	sum := 0

	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-channel // суммируем все значения из канала | 1 + 4 + 9 + 16 + 25 = 55
		}
		fmt.Println(sum) // выводим сумму
		quit <- 0        // посылаем сигнал закрытия
	}()

	sumOfSquare(channel, quit) // запускаем функцию
}
