/*
	Пермутация — это возможное расположение элементов множества в определённом порядке.
	Для заданного среза или строки необходимо сгенерировать все возможные уникальные перестановки её символов.
*/

package main

// permute генерирует все пермутации заданной строки
func permute(str string) []string {
	var result []string

	// функция для рекурсивной генерации пермутаций
	var recurse func(string, string)
	recurse = func(path, available string) {
		if len(available) == 0 {
			result = append(result, path)
			return
		}

		for i := range available {
			recurse(path+string(available[i]), available[:i]+available[i+1:])
		}
	}

	recurse("", str)
	return result
}

func main() {
	str := "abc"
	result := permute(str)
	for _, v := range result {
		println(v)
	}
}
