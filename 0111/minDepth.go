package _111

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//null不参与比较
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right)))) + 1
	}
}
