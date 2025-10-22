package main

// ===========================================================
// Задача 7
// 1. Что будет содержать s после инициализации?
// 2. Что произойдет в println для слайса и для мапы?
// ===========================================================
func a(s []int) {
	s = append(s, 37) // 0 0 0 37 len = 3 + 1 cap = 6
	// s новый слайс!
}

func b(m map[int]int) {
	m[3] = 33
}

func main() {
	s := make([]int, 3, 8) // 0 0 0 len = 3 cap = 3
	m := make(map[int]int, 8)

	// add to slice
	a(s)          // 0 0 0
	println(s[3]) // panic!

	// add to map
	b(m)
	println(m[3]) // 33
}
