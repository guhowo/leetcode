package _053

import "math"

func maxSubArray(nums []int) int {
	states := make([]int, len(nums))
	states[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		states[i] = int(math.Max(float64(states[i-1]+nums[i]), float64(nums[i])))
	}

	max := states[0]
	for i := 0; i < len(nums); i++ {
		if max < states[i] {
			max = states[i]
		}
	}
	return max
}

//2
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = Max(dp[i-1]+nums[i], nums[i])
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
