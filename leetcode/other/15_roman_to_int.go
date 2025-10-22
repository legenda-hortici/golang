package other

func RomanToInt(s string) int {
	symbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := make([]int, 0)
	resultInt := 0
	for i := 0; i < len(s); i++ {
		value, ok := symbols[string(s[i])]
		if !ok {
			return 0
		}
		res = append(res, value)
	}

	for i := 0; i < len(res); i++ {
		temp := 0
		if i < len(res)-1 && res[i] < res[i+1] {
			temp = res[i+1] - res[i]
			i++
		} else {
			temp = res[i]
		}
		resultInt += temp
	}
	return resultInt
}
