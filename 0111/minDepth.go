package _111

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
		l := minDepth(root.Left)
		r := minDepth(root.Right)
		if l < r {
			return l + 1
		} else {
			return r + 1
		}
	}
}
