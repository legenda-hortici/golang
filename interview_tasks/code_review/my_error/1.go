package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "my error"
}

func main() {
	fmt.Println(returnErr() == nil)            // true
	fmt.Println(returnErrorPtr() == nil)       // true
	fmt.Println(returnCustomError() == nil)    // false
	fmt.Println(returnCustomErrorPtr() == nil) // false
	fmt.Println(returnMyError() == nil)        // true

	fmt.Println(returnErr())            // nil
	fmt.Println(returnErrorPtr())       // nil
	fmt.Println(returnCustomError())    // my error
	fmt.Println(returnCustomErrorPtr()) // my error
	fmt.Println(returnMyError())        // my error
}

func returnErr() error {
	var err error
	return err
}

func returnErrorPtr() *error {
	var err *error
	return err
}

func returnCustomError() error {
	var customErr MyError
	return &customErr
}

func returnCustomErrorPtr() error {
	var customErr *MyError
	return customErr
}

func returnMyError() *MyError {
	return nil
}
