package main

import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	if columnNumber <= 26 {
		return string(rune(columnNumber + 64))
	}

	first := columnNumber / 26
	last := columnNumber % 26

	if last == 0 { // если остаток ноль → это Z
		last = 26
		first-- // уменьшаем "первую часть"
	}

	return convertToTitle(first) + string(rune(last+64))
}

func main() {
	fmt.Println(convertToTitle(2147483647))
}
