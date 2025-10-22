package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Дан список целых чисел, повторяющихся элементов в списке нет.
Нужно преобразовать это множество в строку,
сворачивая соседние по числовому ряду числа в диапазоны.

Примеры:
- [1, 4, 5, 2, 3, 9, 8, 11, 0] => "8-9,0-5,11"
- [1, 4, 3, 2] => "1-4"
- [1, 4] => "1,4"
*/

func convert(slice []int) string {
	if len(slice) == 0 {
		return ""
	}

	sort.Ints(slice) // 0 1 2 3 4 5 8 9 11

	var result []string

	start := slice[0] // отмечаем начало диапазона
	end := slice[0]   // отмечаем конец диапазона

	for i := 1; i < len(slice); i++ {
		if slice[i] == end+1 { // сравниваем текущее значение слайса и конец
			end = slice[i] // если равны то перезаписываем конец
		} else { // если значение слайса != концу
			if start == end { // сравниваем начало и конец диапазона
				result = append(result, strconv.Itoa(start)) // если равны то это число и есть диапазон
			} else { // иначе это последовательность
				result = append(result, fmt.Sprintf("%d-%d", start, end))
			}
			start = slice[i] // для следующих значений перезаписываем начало
			end = slice[i]   // и конец
		}
	}

	// когда прошли весь цикл то сравниваем начало и конец для последних чисел слайса
	if start == end { // если равны значит это 1 число
		result = append(result, strconv.Itoa(start))
	} else { // иначе последний диапазон
		result = append(result, fmt.Sprintf("%d-%d", start, end))
	}

	return strings.Join(result, ", ")
}

func main() {
	fmt.Println(convert([]int{1, 4, 5, 2, 3, 9, 8, 11, 0}))
}
