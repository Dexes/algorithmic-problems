package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(first *TreeNode, second *TreeNode) *TreeNode {
	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	first.Val += second.Val
	first.Left = mergeTrees(first.Left, second.Left)
	first.Right = mergeTrees(first.Right, second.Right)

	return first
}
