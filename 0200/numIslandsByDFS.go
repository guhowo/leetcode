package _200

func numIslandsByDFS(grid [][]byte) int {
	result := 0

	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result += recurDFS(i, j, grid, visited)
			}
		}
	}
	return result
}

func recurDFS(x, y int, grid [][]byte, visited [][]bool) int {
	if visited[x][y] {
		return 0
	}

	visited[x][y] = true

	//向上
	if x-1 >= 0 && !visited[x-1][y] && grid[x-1][y] == '1' {
		recurDFS(x-1, y, grid, visited)
	}

	//向下
	if x+1 <= len(grid)-1 && !visited[x+1][y] && grid[x+1][y] == '1' {
		recurDFS(x+1, y, grid, visited)
	}

	//向左
	if y-1 >= 0 && !visited[x][y-1] && grid[x][y-1] == '1' {
		recurDFS(x, y-1, grid, visited)
	}

	//向右
	if y+1 <= len(grid[0])-1 && !visited[x][y+1] && grid[x][y+1] == '1' {
		recurDFS(x, y+1, grid, visited)
	}

	return 1
}
