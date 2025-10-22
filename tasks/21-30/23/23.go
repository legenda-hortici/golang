package main

import "fmt"

func main() {
	i1 := 10 // i1 = 10
	k := 20  // k = 20
	i2 := &k // i2 = 0x00...

	defer fmt.Println(i1) // 10

	defer fmt.Println(*i2) // 20

	defer func(i *int) {
		fmt.Println(*i) // 2020
	}(i2)

	i1 = 1010  // i1 = 1010
	*i2 = 2020 // i2 = 2020
}
