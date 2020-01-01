package _105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	//value -> index
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	return buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inorderMap)
}

func buildTreeHelper(preorder []int, preBegin, preEnd int, inorder []int, inBegin, inEnd int, m map[int]int) *TreeNode {
	if preBegin > preEnd {
		return nil
	}

	root := &TreeNode{preorder[preBegin], nil, nil}
	rootIndex := m[preorder[preBegin]]
	leftNum := rootIndex - inBegin
	root.Left = buildTreeHelper(preorder, preBegin+1, preBegin+leftNum, inorder, inBegin, rootIndex-1, m)
	root.Right = buildTreeHelper(preorder, preBegin+leftNum+1, preEnd, inorder, rootIndex+1, inEnd, m)

	return root
}
