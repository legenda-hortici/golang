package main

import (
	"fmt"
	"sync"
)

//===========================================================
//Задача 12
//1. Конурентно по батчам запросить данные и записать в файл. Нужна общая конструкция, функции которые делают запрос к сайту и выгрузку в файл можно не реализовывать.
//2. Сделать так, чтобы одновременно выполнялось не более chunkSize запросов.
//===========================================================

const url = `http://jsonplaceholder.typicode.com/tools/%d`
const chunkSize = 100
const dataCount = 2 << 10

// имитация запроса
func fetchData(id int) string {
	return fmt.Sprintf("data-%d", id)
}

// имитация записи в файл
func writeToFile(data string) {
	fmt.Println("saved:", data)
}

func main() {
	// канал с задачами
	jobs := make(chan int, dataCount)

	var wg sync.WaitGroup
	for w := 1; w <= chunkSize; w++ { // запуск воркеров кол-вом chunkSize
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range jobs { // берем id задачи
				data := fetchData(id) // передаем в функцию обработчик
				writeToFile(data)     // записываем в файл
			}
		}()
	}

	// кладем все задания
	for i := 1; i <= dataCount; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	fmt.Println("done")
}
