package _104

import "math"

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归版本
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}

//非递归版本
func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	depth := 0
	levelSize := 0
	for len(queue) != 0 {
		depth++
		levelSize = len(queue)
		for i := 0; i < levelSize; i++ {
			p := queue[0]
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
			queue = queue[1:]
		}
	}
	return depth
}
