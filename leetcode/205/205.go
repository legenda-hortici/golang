package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	if len(t) != len(s) {
		return false
	}

	for i := 0; i < len(s); i++ {
		sc, tc := s[i], t[i]

		if val, ok := sMap[sc]; ok {
			if val != int(tc) {
				return false
			}
		} else {
			sMap[sc] = int(tc)
		}

		if val, ok := tMap[tc]; ok {
			if val != int(sc) {
				return false
			}
		} else {
			tMap[tc] = int(sc)
		}

	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}
