package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	return dfs(root, 1)
}

func dfs(root *TreeNode, length int) int {
	subLength, next, result := 0, root.Val+1, length

	if root.Right != nil {
		if subLength = 1; root.Right.Val == next {
			subLength += length
		}

		result = max(result, dfs(root.Right, subLength))
	}

	if root.Left != nil {
		if subLength = 1; root.Left.Val == next {
			subLength += length
		}

		result = max(result, dfs(root.Left, subLength))
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
