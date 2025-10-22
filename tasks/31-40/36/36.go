package main

import (
	"fmt"
	"sort"
)

// Напишите тело функции
func merge(list1, list2 []int) []int {
	// list1, list2 остортированы по возрастанию
	// результат тоже должен быть отсортирован по возрастанию
	n1, n2 := len(list1), len(list2)
	if n1 == 0 && n2 == 0 {
		return []int{}
	} else if n1 == 0 {
		return list2
	} else if n2 == 0 {
		return list1
	}

	for i := 0; i < n1; i++ {
		list1 = append(list1, list2[i])
	}
	sort.Ints(list1)
	return list1
}

func main() {
	fmt.Println(merge([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
}
