package other

func MaxProfit(prices []int) int {
	min := prices[0]
	max := 0
	for i := 0; i < len(prices); i++ {
		if max < prices[i]-min {
			max = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}

	return max
}
