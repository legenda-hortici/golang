package other

func IsValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			if c == ')' && stack[len(stack)-1] == '(' ||
				c == ']' && stack[len(stack)-1] == '[' ||
				c == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
