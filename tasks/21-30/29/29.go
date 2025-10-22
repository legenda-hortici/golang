package main

import "fmt"

// Вернуть из функции ошибку, не подключая доп.пакетов

type myError string

func (e myError) Error() string {
	return string(e)
}

func main() {
	fmt.Println(handle())
}

func handle() error {
	return myError("my error")
}
