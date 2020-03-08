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

//2
func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	diff := make([]int, len(prices)-1)
	for i:=1; i<len(prices); i++ {
		diff[i-1] = prices[i] - prices[i-1]
	}

	dp := make([]int, len(diff))
	dp[0] = diff[0]
	profit := diff[0]
	for i:=1; i<len(diff); i++ {
		dp[i] = Max(dp[i-1]+diff[i], diff[i])
		if profit < dp[i] {
			profit = dp[i]
		}
	}
	if profit < 0 {
		return 0
	}
	return profit
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
