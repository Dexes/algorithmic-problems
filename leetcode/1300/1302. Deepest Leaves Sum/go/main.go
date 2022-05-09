package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	return sum(root, 0, depth(root))
}

func sum(root *TreeNode, currentDepth, maxDepth int) int {
	if root == nil {
		return 0
	}

	if currentDepth == maxDepth {
		return root.Val
	}

	return sum(root.Left, currentDepth+1, maxDepth) + sum(root.Right, currentDepth+1, maxDepth)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
