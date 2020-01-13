package _084

//解法二：单调栈
//核心是维护一个单调递增栈，stack中存放该元素对应的index
//每一个元素出栈意味着以该元素为[顶]的矩形面积计算出来了。
//如果stack还有元素，求以每条边为左边的最大矩阵，此时压入栈中的height为0可以保证栈中每条边都会出栈，因此最后mock一个0入栈
func largestRectangleAreaByStack(h []int) int {
	return 0
}
