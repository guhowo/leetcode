package _122

func maxProfit(prices []int) int {
	if len(prices)<=1 {
		return 0
	}

	profit := 0
	for i:=0; i<len(prices)-1; i++ {
		if  prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}

	return profit
}
