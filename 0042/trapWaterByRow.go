package _042

//方法一：计算每一层能容纳的水
//时间复杂度：O(m*n)
func trapByRow(height []int) int {
	//先求出最高高度
	highest := 0
	for _, h := range height {
		if highest < h {
			highest = h
		}
	}

	volume := 0
	//遍历每次层，从第0层开始
	for h := 1; h <= highest; h++ {
		for i := 0; i < len(height); i++ {
			//第i个柱子比层高要高，则不能储水
			if height[i] >= h {
				continue
			} else {
				//这一格子能积水的条件：i的左右都有>=h的柱子
				if isOK(i, h, height) {
					volume++
				}
			}
		}
	}

	return volume
}

func isOK(pos, h int, height []int) bool {
	l, r := 0, 0
	for l = pos - 1; l >= 0; l-- {
		if height[l] >= h {
			break
		}
	}
	if l == -1 {
		return false
	}

	for r = pos + 1; r < len(height); r++ {
		if height[r] >= h {
			break
		}
	}
	if r == len(height) {
		return false
	}

	return true
}
