package _322

//dp[i]=x表示当目标金额为 i 时，至少需要 x 枚硬币;
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	dp := make([]int, amount+1)
	//因为是求最小值，再将dp初始化为极大值，此处省略
	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = min(dp[i-coins[j]]+1, dp[i])
		}
	}
}

/*
var minCount int
func CoinChange(coins []int, amount int) int {
	minCount = math.MaxInt8
	dfs(coins, 0, 0, amount, 0)
	return minCount
}

func dfs(coins []int, curAmount int, i int, amount int, count int) {
	if i == len(coins) {
		return
	}
	if curAmount > amount{
		return
	}
	if curAmount == amount{
		fmt.Println(i, count)
		minCount = int(math.Min(float64(minCount), float64(count)))
		return
	}
	for idx := 0; idx <= amount / coins[i]; idx++{
		//if curAmount + coins[i] == amount {
		//	fmt.Println(curAmount+coins[i]*idx, amount, i, idx)
		//}
		dfs(coins, curAmount + coins[i] * idx, i + 1, amount, count + idx)
	}
}
*/
