package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	return _countNodes(root, -1, -1)
}

func _countNodes(root *TreeNode, leftHeight, rightHeight int) int {
	if root == nil {
		return 0
	}

	if leftHeight == -1 {
		leftHeight = 1
		for node := root.Left; node != nil; node = node.Left {
			leftHeight++
		}
	}

	if rightHeight == -1 {
		rightHeight = 1
		for node := root.Right; node != nil; node = node.Right {
			rightHeight++
		}
	}

	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	}

	return _countNodes(root.Left, leftHeight-1, -1) + _countNodes(root.Right, -1, rightHeight-1) + 1
}
