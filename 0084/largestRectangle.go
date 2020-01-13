package _084

func largestRectangleArea(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		area := getArea(i, heights)
		if max < area {
			max = area
		}
	}

	return max
}

func getArea(pos int, h []int) int {
	left, right := pos, pos

	for left >= 0 {
		if h[left] < h[pos] {
			break
		}
		left--
	}

	for right < len(h) {
		if h[right] < h[pos] {
			break
		}
		right++
	}

	return (right - left - 1) * h[pos]
}
