package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return getSum(root, 0)
}

func getSum(root *TreeNode, number int) int {
	if root == nil {
		return 0
	}

	number = number*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return number
	}

	return getSum(root.Left, number) + getSum(root.Right, number)
}
