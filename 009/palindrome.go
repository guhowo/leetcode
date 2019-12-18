package _09

func isPalindrome(x int) bool {
	if x<0 {
		return false
	}
	if x == 0{
		return true
	}
	nums := make([]int,0)
	for x !=0 {
		nums = append(nums, x%10)
		x = x/10
	}
	for i,j:=0, len(nums)-1; i<=j; {
		if nums[i]!=nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromeV2(x int) bool {
	if x<0 {
		return false
	}
	if x == 0{
		return true
	}
	shang := 1
	y := x
	for y !=0 {
		shang *= 10
		y = y/10
	}
	shang /=10
	for shang!= 1 {
		if x/shang != x%10 {
			return false
		}
		shang /=100
		x = (x/10 - x%10)/10
	}
	return true
}
