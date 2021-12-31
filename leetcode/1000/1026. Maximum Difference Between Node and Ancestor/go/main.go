package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return getDiff(root, root.Val, root.Val)
}

func getDiff(node *TreeNode, min, max int) int {
	if node == nil {
		return max - min
	}

	min, max = getMin(min, node.Val), getMax(max, node.Val)
	return getMax(getDiff(node.Left, min, max), getDiff(node.Right, min, max))
}

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
