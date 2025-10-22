// Реализация кучи

package main

import "fmt"

// MaxHeap структура кучи
type MaxHeap struct {
	array []int
}

// Insert вставляет элемент в кучу
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) ExtractMax() int {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}

	max := h.array[0]
	// Перемещаем последний элемент в корень
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	h.maxHeapifyDown(0)

	return max
}

// Peek возвращает максимальный элемент в куче без его удаления
func (h *MaxHeap) Peek() int {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}
	return h.array[0]
}

// maxHeapifyUp перемещает элемент вверх в кучу
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[index] > h.array[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}

// maxHeapifyDown перемещает элемент вниз в кучу
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r, largest := left(index), right(index), index
	if l <= lastIndex && h.array[l] > h.array[index] {
		largest = l
	}
	if r <= lastIndex && h.array[r] > h.array[largest] {
		largest = r
	}
	if largest != index {
		h.swap(index, largest)
		h.maxHeapifyDown(largest)
	}
}

// Вспомогательные функции для получение индекса родителя, левого и правого потомка
func parent(index int) int { return (index - 1) / 2 }
func left(index int) int   { return index*2 + 1 }
func right(index int) int  { return index*2 + 2 }

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	heap := &MaxHeap{}
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(15)
	heap.Insert(5)
	heap.Insert(4)
	heap.Insert(45)
	fmt.Println(heap.ExtractMax()) // 45
	fmt.Println(heap.Peek())       // 15
	fmt.Println(heap.ExtractMax()) // 15
}
