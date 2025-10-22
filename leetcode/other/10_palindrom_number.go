package other

import "strconv"

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	for i := 0; i < len(xStr); i++ {
		if xStr[i] != xStr[len(xStr)-i-1] {
			return false
		}
	}
	return true
}
