package _200

func numIslandsByBFS(grid [][]byte) int {
	result := 0

	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	type Node struct {
		x, y int
	}

	queue := make([]Node, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				visited[i][j] = true
				queue = append(queue, Node{i, j})
				for len(queue) != 0 {
					size := len(queue)
					for k := 0; k < size; k++ {
						x := queue[0].x
						y := queue[0].y
						//向上
						if x-1 >= 0 && !visited[x-1][y] && grid[x-1][y] == '1' {
							visited[x-1][y] = true
							queue = append(queue, Node{x - 1, y})
						}
						//向下
						if x+1 <= len(grid)-1 && !visited[x+1][y] && grid[x+1][y] == '1' {
							visited[x+1][y] = true
							queue = append(queue, Node{x + 1, y})
						}
						//向左
						if y-1 >= 0 && !visited[x][y-1] && grid[x][y-1] == '1' {
							visited[x][y-1] = true
							queue = append(queue, Node{x, y - 1})
						}
						//向右
						if y+1 <= len(grid[0])-1 && !visited[x][y+1] && grid[x][y+1] == '1' {
							visited[x][y+1] = true
							queue = append(queue, Node{x, y + 1})
						}
					}
					queue = queue[1:]
				}
				result++
			}
		}
	}
	return result
}
