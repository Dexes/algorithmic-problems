package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}

	return inorder(root, nil, []*TreeNode{nil, nil})
}

func inorder(root *TreeNode, first *TreeNode, window []*TreeNode) *TreeNode {
	if root == nil {
		return first
	}

	first = inorder(root.Left, first, window)

	window[0], window[1] = window[1], root
	if window[0] != nil {
		if first == nil {
			first = window[0]
		}

		window[0].Right, window[1].Left = window[1], nil
	}

	return inorder(root.Right, first, window)
}
