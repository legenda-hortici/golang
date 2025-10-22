package main

import (
	"fmt"
	"math/rand"
)

//===========================================================
//Задача 2
//1. Нужно написать функцию генератор паролей, которая принимает целое число n, а на выходе строка длины n из букв a-zA-Z и 0-9
//2. Что тут можно улучшить?
//3. Какие тесты ты бы написал для нее? Есть какая-нибудь возможность угадать, какая строка будет генерироваться, чтобы писать тесты?
//===========================================================

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generatePassword(n int) string {
	b := make([]byte, n)
	for char := range b {
		b[char] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	fmt.Println(generatePassword(8))
}
