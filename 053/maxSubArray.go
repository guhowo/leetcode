package _53

import "math"

func maxSubArray(nums []int) int {
	states := make([]int, len(nums))
	states[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		states[i] = int(math.Max(float64(states[i-1]+nums[i]), float64(nums[i])))
	}

	max := states[0]
	for i := 0; i < len(nums); i++ {
		if max < states[i] {
			max = states[i]
		}
	}
	return max
}
