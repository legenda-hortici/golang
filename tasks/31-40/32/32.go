package main

import (
	"fmt"
)

/*

// Мы хотим складывать очень большие числа, которые превышают емкость базовых типов, поэтому мы храним их в виде массива неотрицательных чисел.
// Нужно написать функцию, которая примет на вход два таких массива, вычислит сумму чисел, представленных массивами, и вернет результат в виде такого же массива

arr1 = [1, 2, 3] # число 123
arr2 = [4, 5, 6] # число 456
// вывод
res = [5, 7, 9] # число 579. Допустим ответ с первым незначимым нулем [0, 5, 7, 9]

//# Пример 2
# ввод
# ввод
arr1 = [5, 4, 4] # число 544
arr2 = [4, 5, 6] # число 456
# вывод
res = [1, 0, 0, 0] # число 1000
*/

func bigNums(slice1 []int, slice2 []int) []int {
	n1, n2 := len(slice1), len(slice2)
	if n1 < n2 {
		pad := make([]int, n1-n2)
		slice1 = append(pad, slice1...)
	} else if n2 < n1 {
		pad := make([]int, n2-n1)
		slice2 = append(pad, slice2...)
	}
	c := 0
	res := make([]int, len(slice1)+1)
	k := len(res) - 1

	for i := len(slice1) - 1; i >= 0; i-- {
		sum := slice1[i] + slice2[i] + c
		res[k] = sum % 10
		c = sum / 10
		k--
	}
	if c > 0 {
		res[k] = c
	} else {
		res = res[1:]
	}

	return res
}

func main() {
	fmt.Println(bigNums([]int{5, 4, 3}, []int{4, 5, 6}))
}
