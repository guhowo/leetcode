package _042

func trap(height []int) int {
	stack := make([]int, len(height))

	for _, h := range height{
		//前面有两个或以上的矩形，且当前高度大与top才能积水
		if len(stack)>=2 {

		}
		stack = append(stack, h)
	}
}
