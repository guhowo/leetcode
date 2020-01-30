package _121

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))

	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i+1] > dp[i+1] {
			dp[i] = prices[i+1]
		} else {
			dp[i] = dp[i+1]
		}
	}

	max, profit := 0, 0
	for i := 0; i < len(dp)-1; i++ {
		profit = dp[i] - prices[i]
		if max < profit {
			max = profit
		}
	}
	return max
}
