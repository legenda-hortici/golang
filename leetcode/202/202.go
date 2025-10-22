package main

import (
	"fmt"
)

func isHappy(n int) bool {
	seen := make(map[int]bool)
	return check(n, seen)

}

func check(n int, seen map[int]bool) bool {
	if n == 1 {
		return true
	}
	if seen[n] {
		return false
	}
	seen[n] = true

	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}

	return check(sum, seen)
}

func main() {
	fmt.Println(isHappy(19))
}
