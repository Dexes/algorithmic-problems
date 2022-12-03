package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	for ; root != nil; root = root.Right {
		if root.Left == nil {
			continue
		}

		insertRight(root.Left, root.Right)
		root.Left, root.Right = nil, root.Left
	}
}

func insertRight(root, child *TreeNode) {
	if child == nil {
		return
	}

	for root.Right != nil {
		root = root.Right
	}

	root.Right = child
}
