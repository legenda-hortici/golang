package main

import "fmt"

/*
Смотри пример:

arr := []int{1, 2, 3, 4, 5}
sub := arr[:2]
sub = append(sub, 100)
fmt.Println(arr)
fmt.Println(sub)

👉 Объясни результат и почему так происходит.
*/

func main() {
	arr := []int{1, 2, 3, 4, 5} // len = 5, cap = 5
	sub := arr[:2]              // 1 2 | len(sub) = 2
	sub = append(sub, 100)      // 1 2 100 | len(sub) = 3
	fmt.Println(arr)            // 1 2 100 4 5
	fmt.Println(sub)            // 1 2 100
}
