package main

import "fmt"

func main() {
	str := "hello world"

	// Форматируем строку, не выводя её
	str = fmt.Sprintf("formated text %s", str)

	// Теперь мы можем использовать formatStr как нужно
	fmt.Println(str)
}
