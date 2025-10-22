package main

import "fmt"

/*
Есть большой слайс строк (например, 1 млн элементов).
Ты берешь sub := bigSlice[:10].
Объясни, почему sub всё ещё держит в памяти весь большой слайс, и как можно "освободить" память (подсказка: append([]T(nil), sub...)).
*/

func main() {
	bigSlice := make([]int, 0)

	for i := range 1000000 {
		bigSlice = append(bigSlice, i)
	}

	sub := bigSlice[:10] // sub хранит вест слайс, так как имеет ссылку на слайс выше, то есть базовый слайс

	newSub := append([]int(nil), sub...) // здесь создаем новый слайс с новыми данными и только с теми, которые получены в sub и в памяти не будет всего остального bigSlice

	fmt.Println(sub, len(sub), cap(sub))

	fmt.Println(newSub, len(newSub), cap(newSub))

	/*
		bigSlice = [0 1 2 3 4 5 6 7 8 9] len = 10 cap = 1055744
		newSub = [0 1 2 3 4 5 6 7 8 9] len = 10 cap = 10
	*/
}
