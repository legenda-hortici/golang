package main

import "fmt"

func cross(a, b []int) []int {
	counter := make(map[int]int) //создаем счетчик
	var result []int             // выходной слайс

	for _, elem := range a { // пробегаемся по 1 слайсу
		if _, ok := counter[elem]; !ok { // проверяем есть ли элемент в счетчике
			counter[elem] = 1 // если нет то ставим 1
		} else {
			counter[elem] += 1 // иначе увеличиваем на 1
		}
	}

	for _, elem := range b { // пробегаемся по 2 слайсу
		if val, ok := counter[elem]; ok && val > 0 { // проверяем, если есть элемент и он больше 0
			counter[elem] -= 1            // то уменьшаем счетчик
			result = append(result, elem) // добавляем элемент в итоговый слайс
		}
	}

	return result
}

func main() {
	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{6, 2, 8, 1}

	fmt.Println("%v", cross(sl1, sl2)) // [2 1]

	sl1 = []int{1, 1, 1}
	sl2 = []int{1, 1, 1, 1}

	fmt.Println("%v", cross(sl1, sl2)) // [1 1 1]
}
