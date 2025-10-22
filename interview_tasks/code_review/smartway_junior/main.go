package main

import "fmt"

func main() {
	// task 1
	type MyInt int
	var x MyInt = 1
	//var y int = x // cannot use x (variable of int type MyInt) as int value in variable declaration
	var z MyInt = 2

	fmt.Println(x, z)

	// task 2
	// var i interface{} = 5
	// y := i.(string) // panic: interface conversion: interface {} is int, not string
	// fmt.Println(y)

	// task 3
	type account struct{ value int }

	s1 := make([]account, 0, 2) // [] len = 0, cap = 2
	s1 = append(s1, account{})  // [account{value:0}] len = 1 cap 2
	s2 := append(s1, account{}) // [account{value:0} account{value:0}] len = 2 cap = 2

	acc := &s2[0]
	acc.value = 10
	fmt.Println(s1, s2) // s1 = [account{value:10}] s2 = [account{value:10} account{}]

	s1 = append(s2, account{}) // [account{value:0} account{value:0} account{value:0}] len = 3 cap = 4
	acc.value += 10
	fmt.Println(s1, s2) // s1 = [account{value:10} account{} account{}] s2 = [account{value:20} account{}]

	// task 4
	_ = map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	// a := &intKeys[1] // invalid operation: cannot take address of intKeys[1] (map index expression of type string)

	// task 5
	var num *int = nil // nil
	var i interface{} = num
	fmt.Println(i == nil) // this comparison is never true -> false

}
