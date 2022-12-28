package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	result, _ := fillLeaves(root, make([][]int, 0, 16))
	return result
}

func fillLeaves(root *TreeNode, leaves [][]int) ([][]int, int) {
	if root == nil {
		return leaves, -1
	}

	leaves, leftIndex := fillLeaves(root.Left, leaves)
	leaves, rightIndex := fillLeaves(root.Right, leaves)
	index := max(leftIndex, rightIndex) + 1

	if len(leaves) == index {
		leaves = append(leaves, make([]int, 0, 8))
	}

	leaves[index] = append(leaves[index], root.Val)
	return leaves, index
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
