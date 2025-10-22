package main

import "fmt"

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	tmp := make([]int, 0)
	for num > 0 {
		digit := num % 10
		tmp = append(tmp, digit)
		num /= 10
	}
	sum := sumOfDigits(tmp)

	if sum < 10 {
		return sum
	}
	num = sum
	return addDigits(num)
}

func sumOfDigits(num []int) int {
	sum := 0
	for _, i := range num {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(addDigits(38))
}
