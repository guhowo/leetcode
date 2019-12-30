package _094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	recurInorder(root, &res)
	return res
}

//递归版本
func recurInorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	recurInorder(root.Left, res)
	*res = append(*res, root.Val)
	recurInorder(root.Right, res)
}

//非递归版
func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	p := root

	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			res = append(res, p.Val)
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return res
}
