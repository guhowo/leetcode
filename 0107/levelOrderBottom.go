package _107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	//result：遍历的返回结果
	result := make([][]int, 0)

	//用于层次遍历的入队
	queue := make([]*TreeNode, 0)

	if root != nil {
		queue = append(queue, root)
	} else {
		return result
	}

	result = append(result, []int{root.Val})

	for len(queue) != 0 {
		//该层的遍历结果
		levelQue := make([]int, 0)
		//遍历该层
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[0]
			if p.Left != nil {
				queue = append(queue, p.Left)
				levelQue = append(levelQue, p.Left.Val)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
				levelQue = append(levelQue, p.Right.Val)
			}
			queue = queue[1:]
		}
		if len(levelQue) != 0 {
			result = append([][]int{levelQue}, result...)
		}
	}
	return result
}
