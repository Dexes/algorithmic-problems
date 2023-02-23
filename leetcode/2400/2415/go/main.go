package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root.Left != nil {
		reverse(false, root.Left, root.Right)
	}

	return root
}

func reverse(flag bool, left, right *TreeNode) {
	if flag = !flag; flag {
		left.Val, right.Val = right.Val, left.Val
	}

	if left.Left == nil || (flag && left.Left.Left == nil) {
		return
	}

	reverse(flag, left.Left, right.Right)
	reverse(flag, left.Right, right.Left)
}
