package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	wToC := make(map[string]string)
	cToW := make(map[string]string)
	sliceStr := strings.Split(s, " ")

	if len(pattern) != len(sliceStr) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		wToC[sliceStr[i]] = string(pattern[i])
		cToW[string(pattern[i])] = sliceStr[i]
	}

	for i := 0; i < len(pattern); i++ {
		char := string(pattern[i]) // a b
		word := sliceStr[i]        // dog cat

		if val, ok := wToC[word]; ok && val != char {
			return false
		}

		if val, ok := cToW[char]; ok && val != word {
			return false
		}

		wToC[word] = char
		cToW[char] = word
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))

}
