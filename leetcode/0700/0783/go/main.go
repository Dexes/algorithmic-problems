package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	return inorder(root, []int{-1, -1})
}

func inorder(node *TreeNode, window []int) int {
	if node == nil {
		return 100000
	}

	result := inorder(node.Left, window)

	window[0], window[1] = window[1], node.Val
	if window[0] >= 0 {
		result = min(result, window[1]-window[0])
	}

	return min(result, inorder(node.Right, window))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
