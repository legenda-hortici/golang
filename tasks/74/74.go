package main

import "fmt"

//===========================================================
//Задача 4
//Что выведет код и почему?
//===========================================================

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	var err *MyError  // type: *MyError, val: nil
	errorHandler(err) // Error: nil
	err = &MyError{}  // type: *MyError, val: &MyError{}
	errorHandler(err) // Error: "MyError!"
}
