package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum(root, 0)
	return root
}

func sum(root *TreeNode, x int) int {
	if root == nil {
		return x
	}

	root.Val += sum(root.Right, x)
	return sum(root.Left, root.Val)
}
