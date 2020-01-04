package _329

import "math"

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	//记录最大长度
	longest := math.MinInt64

	//二维dp记录以这个点起始的最大长度，初始值为0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			currentLongest := recurDFS(i, j, matrix, dp)
			if longest < currentLongest {
				longest = currentLongest
			}
		}
	}

	return longest
}

func recurDFS(x, y int, matrix [][]int, dp [][]int) int {
	if dp[x][y] != 0 {
		return dp[x][y]
	}

	curLongest := 1

	//向上
	if x-1 >= 0 && matrix[x-1][y] > matrix[x][y] {
		curLongest = recurDFS(x-1, y, matrix, dp) + 1
	}

	//向下
	if x+1 <= len(matrix)-1 && matrix[x+1][y] > matrix[x][y] {
		l := recurDFS(x+1, y, matrix, dp) + 1
		if curLongest < l {
			curLongest = l
		}
	}

	//向左
	if y-1 >= 0 && matrix[x][y-1] > matrix[x][y] {
		l := recurDFS(x, y-1, matrix, dp) + 1
		if curLongest < l {
			curLongest = l
		}
	}

	//向右
	if y+1 <= len(matrix[0])-1 && matrix[x][y+1] > matrix[x][y] {
		l := recurDFS(x, y+1, matrix, dp) + 1
		if curLongest < l {
			curLongest = l
		}
	}

	dp[x][y] = curLongest

	return curLongest
}
