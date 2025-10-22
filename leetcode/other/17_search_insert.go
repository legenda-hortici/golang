package other

/*
Если задан отсортированный массив целых чисел и целевое значение, верните
индекс, если цель найдена. Если нет, верните индекс, на котором она находилась
бы, если бы была вставлена по порядку.
Вы должны написать алгоритм со сложностью O(log n) во времени выполнения.
*/

// runtime: 0ms, mem: 4,83 MB

func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	pos := 0
	for i := 0; i < len(nums); i++ {
		if target > nums[i] {
			pos = i + 1
		} else {
			break
		}
	}
	return pos
}
