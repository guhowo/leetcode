package _300

import (
	"math"
)

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	//初始化dp为1
	dp := make([]int, len(nums))
	longest := 1

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
				if longest < dp[i] {
					longest = dp[i]
				}
			}
		}
	}
	return longest
}
