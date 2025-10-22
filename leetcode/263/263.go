package main

import "fmt"

func isUgly(n int) bool {
	count := make(map[int]int)
	for i := 1; i <= n; i++ {
		if n%i == 0 && i != 2 && i != 3 && i != 5 {
			count[n]++
		}
	}
	fmt.Println(count)
	if count[n] >= 2 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isUgly(1))
}
