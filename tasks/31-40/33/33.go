package main

import "fmt"

/*
// Дан массив целых чисел nums и целое число k. Нужно написать функцию, которая вынимает из массива k наиболее часто встречающихся элементов.
// Пример
// ввод
nums = [1,1,1,2,2,3]
k = 2
// вывод (в любом порядке)
[1, 2]
*/

func findPeriod(nums []int, k int) []int {
	check := make(map[int]int)
	res := make([]int, 0)

	seen := 0
	for i := 0; i < len(nums); i++ {
		check[nums[i]]++
		if check[nums[i]] >= k && seen != nums[i] {
			seen = nums[i]
			res = append(res, nums[i])
		}
	}

	fmt.Println(check)

	return res
}

func main() {
	fmt.Println(findPeriod([]int{1, 1, 1, 2, 2, 3}, 2))
}
