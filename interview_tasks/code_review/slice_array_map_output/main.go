package main

import "fmt"

func modifySlice(s []string) {
	s[0] = "слайс"
	s[1] = "модифицирован"
}

func modifyArray(a [2]string) {
	a[0] = "массив"
	a[1] = "модифицирован"
}

func mapModify(m map[int]string) {
	m[0] = "мапа"
	m[1] = "модифицирована"
}

func main() {
	slice := []string{"строка1", "строка2"}
	array := [2]string{"строка1", "строка2"}
	mapSlice := map[int]string{0: "строка1", 1: "строка2"}

	modifySlice(slice)
	fmt.Println(slice) // [слайс модифицирован]

	modifyArray(array)
	fmt.Println(array) // [строка1 строка2]

	mapModify(mapSlice)
	fmt.Println(mapSlice) // [мапа модифицирована]
}
