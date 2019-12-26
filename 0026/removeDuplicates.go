package _026

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	length := len(nums)
	last := length - 1
	moved := 0
	i := length - 2
	for i >= 0 {
		if nums[i] != nums[last] {
			last = i
			i--
			continue
		} else {
			//计算步长
			for i >= 0 && nums[i] == nums[last] {
				i--
			}
			moved = last - i - 1
			//移动
			for k := last; k < length; k++ {
				nums[k-moved] = nums[k]
			}
			length -= moved
			last = i
		}
	}

	return length
}
