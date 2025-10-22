package main

import (
	"fmt"
	"reflect"
)

// ===========================================================
// Задача 2
// 1. Добавить код, который выведет тип переменной whoami
// ===========================================================
func printType(whoami interface{}) {
	val := reflect.TypeOf(whoami)
	fmt.Println("type:", val)
}

func main() {
	printType(42)
	printType("im string")
	printType(true)
}
