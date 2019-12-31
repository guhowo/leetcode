package _102

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([][]int, 0)

	level, levelSize := 0, 0

	for len(queue) != 0 {
		levelSize = len(queue)
		tmp := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			p := queue[0]
			tmp[i] = p.Val
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
			queue = queue[1:]
			level++
		}
		result = append(result, tmp)
	}
	return result
}
