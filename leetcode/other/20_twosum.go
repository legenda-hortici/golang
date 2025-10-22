package other

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		need := target - num
		if idx, found := seen[need]; found {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}
