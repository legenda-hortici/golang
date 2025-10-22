package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [10]int{} // объявили массив на 10 элементов
	resultChan1 := make(chan int, 1)
	resultChan2 := make(chan int, 1)

	for i := 0; i < 10; i++ {
		array[i] = i // заполняем массив
	}

	fmt.Println(array)

	part1 := array[len(array)/2:] // делим массив
	part2 := array[:len(array)/2]

	fmt.Println(part1, part2)

	var wg sync.WaitGroup
	wg.Add(2) // добавляем 2 горутины в очередь

	go func() { // выполняем первую горутину и суммируем все числа из 1 части
		defer wg.Done()
		result := 0
		for _, i := range part1 {
			result += i
			fmt.Printf("goroutine 1: %d\n", result)
		}
		resultChan1 <- result // записываем в 1 канал результат суммы
	}()

	go func() { // выполняем вторую горутину и суммируем все числа из 2 части
		defer wg.Done()
		result := 0
		for _, i := range part2 {
			result += i
			fmt.Printf("goroutine 2: %d\n", result)
		}
		resultChan2 <- result // записываем во 2 канал результат суммы
	}()

	fmt.Println(<-resultChan1 + <-resultChan2) // выводим сумму всех чисел массива

	wg.Wait() // ожидаем выполнения всех горутин и блокируем основной поток
}
