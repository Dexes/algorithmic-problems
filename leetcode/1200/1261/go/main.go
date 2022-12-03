package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements map[int]struct{}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0

	result := make(FindElements)
	fill(root, result)

	return result
}

func (x FindElements) Find(target int) bool {
	_, result := x[target]
	return result
}

func fill(root *TreeNode, set map[int]struct{}) {
	set[root.Val] = struct{}{}

	if root.Left != nil {
		root.Left.Val = root.Val*2 + 1
		fill(root.Left, set)
	}

	if root.Right != nil {
		root.Right.Val = root.Val*2 + 2
		fill(root.Right, set)
	}
}
