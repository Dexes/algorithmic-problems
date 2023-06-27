package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(first, second *TreeNode, target int) bool {
	if first == nil {
		return false
	}

	return search(second, target-first.Val) ||
		twoSumBSTs(first.Left, second, target) ||
		twoSumBSTs(first.Right, second, target)
}

func search(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Val > target {
		return search(root.Left, target)
	}

	if root.Val < target {
		return search(root.Right, target)
	}

	return true
}
