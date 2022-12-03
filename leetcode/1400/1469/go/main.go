package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLonelyNodes(root *TreeNode) []int {
	return search(root, make([]int, 0, 16))
}

func search(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}

	if root.Left != nil && root.Right == nil {
		result = append(result, root.Left.Val)
		return search(root.Left, result)
	}

	if root.Left == nil && root.Right != nil {
		result = append(result, root.Right.Val)
		return search(root.Right, result)
	}

	result = search(root.Left, result)
	result = search(root.Right, result)
	return result
}
