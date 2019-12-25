package _55

func canJump(nums []int) bool {
	//return btSolveJump(0, len(nums), nums)
	return solveJump(0, len(nums), nums)
}

func solveJump(begin, n int, nums []int) bool {
	if begin == n-1 {
		return true
	}
	if nums[begin] == 0 {
		return false
	}

	largestPosition := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 && largestPosition <= i && i != n-1 {
			return false
		}
		if largestPosition < nums[i]+i {
			largestPosition = nums[i] + i
		}
	}
	return true
}

// backtrace
func btSolveJump(i, n int, nums []int) bool {
	if i == n-1 {
		return true
	}
	if nums[i] == 0 {
		return false
	}

	for step := 1; step <= nums[i]; step++ {
		if !btSolveJump(i+step, n, nums) {
			continue
		} else {
			return true
		}
	}
	return false
}
