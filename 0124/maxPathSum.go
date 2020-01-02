package _124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := math.MinInt64

	helper(root, &max)
	return max
}

//以root为根的最大路径可能是root.val
// 或者root.val+以root.left为根的最大距离
// 或者root.val+以root.right为根的最大距离
func helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	//记录以root为根的最大值，暂时以root->val为最大
	cur := root.Val
	leftMax := helper(root.Left, max)
	rightMax := helper(root.Right, max)

	//判断以当前结点为根,是不是最大值
	if leftMax > 0 {
		cur += leftMax
	}
	if rightMax > 0 {
		cur += rightMax
	}
	//记录最大值
	*max = int(math.Max(float64(cur), float64(*max)))
	//但是当前的最大值不一定是以root为根的，sum函数返回的是以root为根的最大值，而不是max
	if leftMax > rightMax && leftMax > 0 {
		return leftMax + root.Val
	} else if rightMax > leftMax && rightMax > 0 {
		return rightMax + root.Val
	} else {
		return root.Val
	}
}
