package _0098

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0) //记录中序遍历的结果

	p := root

	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			if len(res) > 0 && p.Val <= res[len(res)-1] {
				return false
			} else {
				res = append(res, p.Val)
			}
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}

	return true
}
