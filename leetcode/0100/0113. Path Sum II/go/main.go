package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	return dfs(root, targetSum, make([]int, 1000), 0, make([][]int, 0, 500))
}

func dfs(root *TreeNode, targetSum int, path []int, depth int, result [][]int) [][]int {
	if root == nil {
		return result
	}

	targetSum -= root.Val
	path[depth] = root.Val

	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			return appendPath(result, path, depth)
		}

		return result
	}

	result = dfs(root.Left, targetSum, path, depth+1, result)
	result = dfs(root.Right, targetSum, path, depth+1, result)

	return result
}

func appendPath(paths [][]int, path []int, depth int) [][]int {
	x := make([]int, depth+1)
	copy(x, path)

	return append(paths, x)
}
