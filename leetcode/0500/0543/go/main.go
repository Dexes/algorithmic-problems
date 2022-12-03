package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := depthAndDiameter(root)
	return diameter
}

func depthAndDiameter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	lDepth, lDiameter := depthAndDiameter(root.Left)
	rDepth, rDiameter := depthAndDiameter(root.Right)
	return max(lDepth, rDepth) + 1, max(max(lDiameter, rDiameter), lDepth+rDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
