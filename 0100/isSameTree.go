package _100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q != nil) {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	} else {
		a := isSameTree(p.Left, q.Left)
		b := isSameTree(p.Right, q.Right)
		if a == false || b == false {
			return false
		} else {
			return true
		}
	}

	return true
}
