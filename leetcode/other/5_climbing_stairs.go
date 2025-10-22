package other

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		temp := first + second
		first = second
		second = temp
	}
	return second
}
