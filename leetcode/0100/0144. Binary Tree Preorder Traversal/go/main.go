package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return dfs(root, make([]int, 0, 100))
}

func dfs(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = dfs(root.Left, result)
	result = dfs(root.Right, result)

	return result
}
