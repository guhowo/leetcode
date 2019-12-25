package _041

func firstMissingPositive(nums []int) int {
	nMap := make(map[int]struct{})
	for _, v := range nums {
		nMap[v] = struct{}{}
	}

	i := 1
	for {
		if _, ok := nMap[i]; !ok {
			return i
		}
		i++
	}
}
