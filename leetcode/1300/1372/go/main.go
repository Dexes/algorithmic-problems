package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	DirectionLeft  = false
	DirectionRight = true
)

func longestZigZag(root *TreeNode) int {
	return dfs(root, DirectionLeft, -1)
}

func dfs(root *TreeNode, direction bool, length int) int {
	if root == nil {
		return length
	}

	var startNode, zigZagNode *TreeNode
	if direction == DirectionLeft {
		startNode, zigZagNode = root.Left, root.Right
	} else {
		startNode, zigZagNode = root.Right, root.Left
	}

	return max(dfs(startNode, direction, 0), dfs(zigZagNode, !direction, length+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
