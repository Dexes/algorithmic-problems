package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	window := []int{-1, -1}
	dfs(root, window)

	return window[1]
}

func dfs(root *TreeNode, window []int) {
	if root == nil {
		return
	}

	dfs(root.Left, window)
	set(window, root.Val)
	dfs(root.Right, window)
}

func set(window []int, x int) {
	if window[0] == -1 || x < window[0] {
		window[0], window[1] = x, window[0]
		return
	}

	if window[0] != x && (window[1] == -1 || x < window[1]) {
		window[1] = x
	}
}
