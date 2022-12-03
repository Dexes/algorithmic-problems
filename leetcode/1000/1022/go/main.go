package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, make([]int, 0, 32))
}

func dfs(root *TreeNode, current []int) int {
	if root == nil {
		return 0
	}

	current = append(current, root.Val)
	if root.Left == nil && root.Right == nil {
		return toInt(current)
	}

	return dfs(root.Left, current) + dfs(root.Right, current)
}

func toInt(digits []int) int {
	result, right := 0, len(digits)-1
	for i := right; i >= 0; i-- {
		result += digits[i] << (right - i)
	}

	return result
}
