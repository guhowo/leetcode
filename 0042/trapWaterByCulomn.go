package _042

//解法二：计算每一列能容纳的水
//时间复杂度：O(n*n)
func trapByColumn(height []int) int {
	volume := 0

	//遍历每一列，计算每一列的积水量
	for i := 0; i < len(height); i++ {
		maxLeft := getMax(0, i-1, height)
		if maxLeft > height[i] {
			maxRight := getMax(i+1, len(height)-1, height)
			if maxRight > height[i] {
				if maxLeft > maxRight {
					volume += maxRight - height[i]
				}
			}
		}
	}

	return volume
}

func getMax(from, to int, h []int) int {
	var max int
	for i := from; i <= to; i++ {
		if max < h[i] {
			max = h[i]
		}
	}
	return max
}
