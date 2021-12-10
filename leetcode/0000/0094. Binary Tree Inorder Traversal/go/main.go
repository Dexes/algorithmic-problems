package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return inorder(root, make([]int, 0, 100))
}

func inorder(root *TreeNode, data []int) []int {
	if root == nil {
		return data
	}

	data = inorder(root.Left, data)
	data = append(data, root.Val)
	data = inorder(root.Right, data)

	return data
}
