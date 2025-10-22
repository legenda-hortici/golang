package main

import "fmt"

//===========================================================
//Задача 5
//1. Как будет работать код?
//2. Как сделать так, чтобы выводился только первый ch?
//===========================================================

func main() {
	/*
		так как каналы небуферизированные, то во время выполнения select будет выбран случайный кейс
		потому что, все 3 канала готовы к чтению. этот select не является lock-free из-за отсутствия
		default кейса, а значит будет блокировка до тех пор, пока не будет готов 1 из кейсов.

		чтобы выводился только 1 канал, можно добавить default кейс или сделать кейс приоритетным:
		select {
		case <-ch:
			fmt.Printf("val from ch\n")
		default:
		case <-ch2:
			fmt.Printf("val from ch2\n")
		case <-ch3:
			fmt.Printf("val from ch3\n")
		}
	*/

	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	go func() {
		ch <- true
	}()
	go func() {
		ch2 <- true
	}()
	go func() {
		ch3 <- true
	}()

	select {
	default:
		fmt.Printf("val from ch\n")
	case <-ch2:
		fmt.Printf("val from ch2\n")
	case <-ch3:
		fmt.Printf("val from ch3\n")
	}

}
