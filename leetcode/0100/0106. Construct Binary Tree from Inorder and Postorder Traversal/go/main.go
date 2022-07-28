package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	lastIndex := len(postorder) - 1
	root := &TreeNode{Val: postorder[lastIndex]}

	i := search(inorder, postorder[lastIndex])
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:lastIndex])

	return root
}

func search(data []int, target int) (result int) {
	for ; data[result] != target; result++ {
	}

	return result
}
