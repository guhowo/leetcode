package _051

//N皇后问题
func solveNQueens(n int) [][]string {
	if n < 0 {
		return nil
	}
	output := make([][]string, 0)
	result := make([]int, n)
	NQueens(0, n, result, &output)
	return output
}

func NQueens(row int, n int, result []int, output *[][]string) {
	if row == n {
		print(result, n, output)
		return
	}
	for i := 0; i < n; i++ {
		if isOk(row, i, n, result) {
			result[row] = i
			NQueens(row+1, n, result, output)
		}
	}
}

func isOk(row, column, n int, result []int) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		//在同一列
		if result[i] == column {
			return false
		}
		if leftup >= 0 {
			if result[i] == leftup {
				return false
			}
			leftup--
		}
		if rightup < n {
			if result[i] == rightup {
				return false
			}
			rightup++
		}
	}
	return true
}
func print(result []int, n int, output *[][]string) {
	out := make([]string, n)
	for i := 0; i < len(result); i++ {
		rowStr := ""
		for j := 0; j < len(result); j++ {
			if result[j] == i {
				rowStr += "Q"
			} else {
				rowStr += "."
			}
			out[i] = rowStr
		}
	}
	*output = append(*output, out)
}
