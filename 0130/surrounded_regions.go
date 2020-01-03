package _130

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	type Node struct {
		row, col int
	}

	//用于存放#的坐标
	queue := make([]Node, 0)

	row, col := len(board)-1, len(board[0])-1
	//扫描边
	//扫描第0行和最后一行
	for i := 0; i <= col; i++ {
		//第一行
		if board[0][i] == 'O' {
			board[0][i] = '#'
			queue = append(queue, Node{0, i})
		}
		//最后一行
		if board[row][i] == 'O' {
			board[row][i] = '#'
			queue = append(queue, Node{row, i})
		}
	}
	//扫描第0列和最后一列
	for i := 0; i <= row; i++ {
		//第一列
		if board[i][0] == 'O' {
			board[i][0] = '#'
			queue = append(queue, Node{i, 0})
		}
		//最后一列
		if board[i][col] == 'O' {
			board[i][col] = '#'
			queue = append(queue, Node{i, col})
		}
	}

	for len(queue) != 0 {
		size := len(queue)
		for idx := 0; idx < size; idx++ {
			q := queue[0]
			row, col := q.row, q.col
			if row-1 > 0 && board[row-1][col] == 'O' {
				board[row-1][col] = '#'
				queue = append(queue, Node{row - 1, col})
			}
			if row+1 < len(board)-1 && board[row+1][col] == 'O' {
				board[row+1][col] = '#'
				queue = append(queue, Node{row + 1, col})
			}
			if col-1 > 0 && board[row][col-1] == 'O' {
				board[row][col-1] = '#'
				queue = append(queue, Node{row, col - 1})
			}
			if col+1 < len(board[0])-1 && board[row][col+1] == 'O' {
				board[row][col+1] = '#'
				queue = append(queue, Node{row, col + 1})
			}
			queue = queue[1:]
		}
	}

	for i := 0; i <= row; i++ {
		for j := 0; j <= col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}
