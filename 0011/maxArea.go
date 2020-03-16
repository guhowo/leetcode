package _011

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i != j {
		tmp := Area(i, j, height)
		if tmp > max {
			max = tmp
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}

func Area(i, j int, height []int) int {
	if height[i] > height[j] {
		return (j - i) * height[j]
	}
	return (j - i) * height[i]
}
