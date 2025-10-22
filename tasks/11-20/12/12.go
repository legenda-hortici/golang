package main

import "fmt"

/*
Дан слайс []string. Верни новый слайс без повторяющихся строк.
👉 Подумай: порядок нужно сохранять или нет?
*/

func main() {
	strs := []string{"asd", "asd", "asd", "qwe", "xcv", "dgf"}

	fmt.Printf("original slice: %s\n", strs)

	mapped := make(map[string]int)

	for i := 0; i < len(strs); i++ {
		mapped[strs[i]]++
	}

	for i := 0; i < len(strs); i++ {
		if mapped[strs[i]] > 1 {
			strs = append(strs[:i], strs[i+1:]...)
		}
	}

	fmt.Printf("slice after deleted dublicates: %s\n", strs)
}
