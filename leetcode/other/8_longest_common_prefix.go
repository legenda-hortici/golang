package other

func LongestCommonPrefix(strs []string) string {
	res := ""

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res
}
