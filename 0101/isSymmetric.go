package _101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	if !helper(left.Left, right.Right) {
		return false
	}
	if !helper(left.Right, right.Left) {
		return false
	}
	return true
}
