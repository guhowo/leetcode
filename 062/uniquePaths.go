package _62

func uniquePaths(m int, n int) int {
	states := make([][]int, m)
	for i := 0; i < m; i++ {
		states[i] = make([]int, n)
	}

	//初始化第0行
	for i := 0; i < n; i++ {
		states[0][i] = 1
	}
	//初始化第0列
	for i := 0; i < m; i++ {
		states[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			states[i][j] = states[i-1][j] + states[i][j-1]
		}
	}

	return states[m-1][n-1]
}
