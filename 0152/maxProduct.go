package _152

func maxProduct(nums []int) int {
	if len(nums) ==0 {
		return 0
	}
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	max := nums[0]

	for i:=1; i<len(nums); i++ {
		dpMax[i] = getMax(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
		dpMin[i] = getMin(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i], nums[i])
		if max < dpMax[i] {
			max = dpMax[i]
		}
	}

	return max
}

func getMax(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}

func getMin(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
