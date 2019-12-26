package _064

import "math"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	states := make([][]int, m)
	for i := 0; i < m; i++ {
		states[i] = make([]int, n)
	}

	states[0][0] = grid[0][0]
	//初始化第0行
	for j := 1; j < n; j++ {
		states[0][j] = states[0][j-1] + grid[0][j]
	}
	//初始化第0列
	for i := 1; i < m; i++ {
		states[i][0] = states[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			states[i][j] = int(math.Min(float64(states[i-1][j]), float64(states[i][j-1]))) + grid[i][j]
		}
	}
	return states[m-1][n-1]
}
