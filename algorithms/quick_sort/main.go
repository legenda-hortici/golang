// Реализация быстрой сортировки

package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1 // Индексы крайнего левого и правого элементов

	// Выбираем опорный элемент
	pivotIndex := len(arr) / 2

	// Перемещаем опорный элемент в конец
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Разделяем на подмассивы меньше и больше опорного
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	// Возвращаем опорный элемент на место
	arr[left], arr[right] = arr[right], arr[left]

	// Рекурсивно сортируем подмассивы
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	sample := []int{9, -3, 5, 2, 6, 8, -6, 1, 3}
	sorted := quickSort(sample)
	fmt.Println("Отсортированный массив:", sorted)
}
