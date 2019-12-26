package _0718

import "math"

func findLength(A []int, B []int) int {
	dp := make([][]int, len(A))
	for i, _ := range dp {
		dp[i] = make([]int, len(B))
	}

	//初始化第一列
	for i := 0; i < len(A); i++ {
		if A[i] == B[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	//初始化第一行
	for j := 0; j < len(B); j++ {
		if A[0] == B[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	longest := 0
	for i := 1; i < len(A); i++ {
		for j := 1; j < len(B); j++ {
			if A[i] == B[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				longest = int(math.Max(float64(dp[i][j]), float64(longest)))
			} else {
				dp[i][j] = 0
			}
		}
	}

	return longest
}
