package main

import "fmt"

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func checkErr(err error) {
	fmt.Println(err == nil)
}

func main() {
	var e1 error
	checkErr(e1) // true: type error, val nil

	var e *errorString // addr
	checkErr(e) // false: type *errorString, val nil

	e = &errorString{}
	checkErr(e) // false: type *errorString, val not nil

	e = nil // type *errorString, val nil
	checkErr(e) // fasle: type *errorString, val nil
}
