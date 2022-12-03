package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, result := solve(root)
	return result
}

func solve(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftBalanced := solve(root.Left)
	if !leftBalanced {
		return 0, false
	}

	rightDepth, rightBalanced := solve(root.Right)
	if !rightBalanced {
		return 0, false
	}

	return 1 + max(leftDepth, rightDepth), abs(leftDepth-rightDepth) <= 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
