package _155

import "math"

type MinStack struct {
	//用于存储值
	Stack []int
	//用于存储出现过的最小值
	Min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Stack: make([]int, 0), Min: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	if len(this.Min) == 0 || (len(this.Min) > 0 && this.Min[len(this.Min)-1] >= x) {
		this.Min = append(this.Min, x)
	}
}

func (this *MinStack) Pop() {
	size := len(this.Stack)
	p := 0
	if size > 0 {
		p = this.Stack[size-1]
		this.Stack = this.Stack[:size-1]
	}
	minSize := len(this.Min)
	if minSize > 0 && this.Min[minSize-1] == p {
		this.Min = this.Min[:minSize-1]
	}
}

func (this *MinStack) Top() int {
	size := len(this.Stack)
	if size > 0 {
		return this.Stack[size-1]
	}
	return math.MinInt64
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
