package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}

	add(root, val, depth-1)
	return root
}

func add(root *TreeNode, val, depth int) {
	if root == nil {
		return
	}

	if depth == 1 {
		root.Left = &TreeNode{Val: val, Left: root.Left}
		root.Right = &TreeNode{Val: val, Right: root.Right}
		return
	}

	add(root.Left, val, depth-1)
	add(root.Right, val, depth-1)
}
