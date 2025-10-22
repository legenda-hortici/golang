package main

import "fmt"

/*
Есть [][]int (слайс слайсов).
Объясни разницу между:

b := a

и

b := make([][]int, len(a))
for i := range a {
    b[i] = append([]int(nil), a[i]...)
}

👉 Почему первый вариант может привести к неожиданным изменениям?
*/

func main() {
	martix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	fmt.Println(martix)

	b := martix // поверхностная копия (shared memory).

	fmt.Println(b)

	c := make([][]int, len(martix)) // глубокая копия (новые данные).
	for i := range martix {
		c[i] = append([]int(nil), martix[i]...)
	}

	c[0][0] = 888
	fmt.Println("martix:", martix)
	fmt.Println("c:", c)

}
