// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

const (
	eng = " fly me to the moon "
	ru  = "  сел в   озере   березов лес"
)

// eng = str | str[i] -

func polindrom(str string) bool {
	runes := []rune(str)
	res := make([]rune, 0)
	for _, r := range runes {
		if r != 32 {
			res = append(res, r)
		}
	}
	fmt.Println(string(res))

	for i := 0; i < len(res)-1; i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(polindrom(eng))
	fmt.Println(polindrom(ru))
}
