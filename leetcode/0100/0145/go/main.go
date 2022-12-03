package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	return dfs(root, make([]int, 0, 100))
}

func dfs(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}

	result = dfs(root.Left, result)
	result = dfs(root.Right, result)
	return append(result, root.Val)
}
