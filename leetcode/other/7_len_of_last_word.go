package other

import (
	"strings"
)

/*
Если задана строка s, состоящая из слов и пробелов, верните
длину последнего слова в строке.
Слово - это максимальная подстрока, состоящая только из символов,
не содержащих пробелов.
*/

// runtime: 0ms, mem: 4,18 MB

func LengthOfLastWord(s string) int {
	word := []string{}
	str := strings.TrimSpace(s)
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) != " " {
			word = append(word, string(str[i]))
		} else {
			break
		}
	}
	return len(word)
}
