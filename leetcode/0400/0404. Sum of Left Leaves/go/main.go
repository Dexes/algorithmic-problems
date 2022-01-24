package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return dfs(root, false)
}

func dfs(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		if isLeft {
			return root.Val
		}

		return 0
	}

	return dfs(root.Left, true) + dfs(root.Right, false)
}
