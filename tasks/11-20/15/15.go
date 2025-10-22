package main

import "fmt"

/*
Есть слайс []int{1,2,3,4,5,6,7,8}.
Раздели его на два слайса: чётные и нечётные числа.
*/

func main() {
	foo := make([]int, 0)
	bar := make([]int, 0)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			foo = append(foo, slice[i])
		} else {
			bar = append(bar, slice[i])
		}
	}

	fmt.Printf("foo: %d\n", foo)
	fmt.Printf("bar: %d\n", bar)
}
