package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	return sum(root, 1, 1)
}

func sum(node *TreeNode, gpValue, pValue int) int {
	if node == nil {
		return 0
	}

	result := sum(node.Left, pValue, node.Val) + sum(node.Right, pValue, node.Val)
	if gpValue&1 == 0 {
		return result + node.Val
	}

	return result
}
