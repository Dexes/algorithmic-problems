package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, root, k)
}

func dfs(root, current *TreeNode, k int) bool {
	if current == nil {
		return false
	}

	difference := k - current.Val
	if difference != current.Val && exists(root, difference) {
		return true
	}

	return dfs(root, current.Left, k) || dfs(root, current.Right, k)
}

func exists(root *TreeNode, x int) bool {
	if root == nil {
		return false
	}

	if x < root.Val {
		return exists(root.Left, x)
	}

	if x > root.Val {
		return exists(root.Right, x)
	}

	return true
}
