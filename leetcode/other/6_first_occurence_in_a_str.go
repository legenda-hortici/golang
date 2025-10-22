package other

import "fmt"

/*
Если даны две строки needle и haystack, верните индекс первого вхождения needle
в haystack или -1, если needle не является частью haystack.
*/

// runtime: 0ms, beats 100%, mem: 3,86 MB

func StrStr(haystack, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	} else if haystack == needle {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if needle != haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(StrStr("hello", "ll"))
}
