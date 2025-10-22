package main

import (
	"bytes"
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	strSplit := strings.Split(s, " ")

	var buf bytes.Buffer
	for _, v := range strSplit {
		buf.WriteString(strings.ToLower(v))
	}

	bufStr := buf.String()

	for i := 0; i < len(bufStr); i++ {
		if (33 <= byte(bufStr[i]) && byte(bufStr[i]) <= 47) ||
			(58 <= byte(bufStr[i]) && byte(bufStr[i]) <= 64) ||
			(91 <= byte(bufStr[i]) && byte(bufStr[i]) <= 96) ||
			(123 <= byte(bufStr[i]) && byte(bufStr[i]) <= 126) {
			bufStr = bufStr[:i] + bufStr[i+1:]
			fmt.Println(bufStr)
			i--
		}
	}

	fmt.Println(bufStr)

	for i := 0; i < len(bufStr); i++ {
		//fmt.Println(string(bufStr[i]), string(bufStr[len(bufStr)-1-i]))
		if bufStr[i] != bufStr[len(bufStr)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("......a....."))
}
