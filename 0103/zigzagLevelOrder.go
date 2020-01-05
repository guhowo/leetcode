package _103

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

func zigzagLevelOrder(root *TreeNode) [][]int {
	//最终返回结果
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	if root == nil {
		return result
	}
	queue = append(queue, root)
	level := 0

	for len(queue) != 0 {
		//存储每一层的val
		que := make([]int, 0)

		//存放该层遍历结果
		for i := 0; i < len(queue); i++ {
			if level%2 == 0 {
				que = append(que, queue[i].Val)
			} else {
				que = append(que, queue[len(queue)-i-1].Val)
			}
		}
		result = append(result, que)

		//遍历本层
		level++
		size := len(queue)
		for i := 0; i < size; i++ {
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
	return result
}
