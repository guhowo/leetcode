package _042

//解法三：DP
//Dp是对根据列来计算积水量的一种优化
//by column中求maxleft和maxright每次都是重复遍历，其实可以用动态规划
//本解法中maxLeft表示第i个（不含i）的左侧最高高度
func trapByDp(height []int) int {
	volume := 0

	//初始化dp []int
	maxLeft, maxRight := make([]int, len(height)), make([]int, len(height))

	for i := 1; i < len(height); i++ {
		if height[i-1] > maxLeft[i-1] {
			maxLeft[i] = height[i-1]
		} else {
			maxLeft[i] = maxLeft[i-1]
		}
	}

	for i := len(height) - 2; i >= 0; i-- {
		if height[i+1] > maxRight[i+1] {
			maxRight[i] = height[i+1]
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}

	//扫描每一列
	for i := 0; i < len(height); i++ {
		if height[i] < maxLeft[i] && height[i] < maxRight[i] {
			if maxLeft[i] > maxRight[i] {
				volume += maxRight[i] - height[i]
			} else {
				volume += maxLeft[i] - height[i]
			}
		}
	}

	return volume
}
