// Реализация сортировки пузырьком и сортировка вставками

package main

import "fmt"

// сортировка пузырьком
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// сортировка вставками
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Перемещаем элементы arr[0..i-1], которые больше чем key, на одну позицию вперед
		// от текущей позиции
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	sample := []int{64, 34, 25, 12, 22, 11, 90}
	sorted := bubbleSort(sample)
	fmt.Println("Отсортированный массив:", sorted)

	sample = []int{93, 57, 23, 45, 31, 22}
	sorted = insertionSort(sample)
	fmt.Println("Отсортированный массив:", sorted)
}
