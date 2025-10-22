package main

import (
	"fmt"
	"strconv"
)

// Дан массив целых чисел, нужно найти непустой подотрезок (непрерывную подпоследовательность)
// с заданной суммой target, либо сказать, что это невозможно.
// find_target([9, -6, 5, 1, 4, -2], 10) -> (2, 4)

func find_target(slice []int, target int) string {
	if len(slice) == 0 {
		return ""
	}

	for i := 0; i < len(slice); i++ { // проходим циклом по слайсу
		sum := 0                          // сумма элементов слайса
		for j := i; j < len(slice); j++ { // проходим сдвигом по слайсу
			sum += slice[j]    // складываем все элементы слайса
			if sum == target { // сравниваем target и sum
				if i == j { // если это последовательность из 1 числа
					return strconv.Itoa(i)
				} // иначе выводим последовательность
				return fmt.Sprintf("(%d-%d)", i, j)
			}
		}
	}

	return "невозможно"
}

func main() {
	fmt.Println(find_target([]int{9, -6, 5, 1, 4, -2}, 10))
	fmt.Println(find_target([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(find_target([]int{5, -2, 7}, 5))
	fmt.Println(find_target([]int{1, 2, 3}, 100))
}
