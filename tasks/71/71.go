package main

import "fmt"

//==========================================
//Задача 1
//Что выведет код?
//==========================================

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil

}
func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	fmt.Println(a == b) // false
}
