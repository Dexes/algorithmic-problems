package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	sums := getSums(root, 0, make([]int, 0, 64))

	root.Val = 0
	replaceValues(root, 1, sums)

	return root
}

func replaceValues(root *TreeNode, depth int, sums []int) {
	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Left == nil {
		root.Right.Val = sums[depth] - root.Right.Val
		replaceValues(root.Right, depth+1, sums)
		return
	}

	if root.Right == nil {
		root.Left.Val = sums[depth] - root.Left.Val
		replaceValues(root.Left, depth+1, sums)
		return
	}

	val := sums[depth] - root.Right.Val - root.Left.Val
	root.Left.Val, root.Right.Val = val, val
	replaceValues(root.Left, depth+1, sums)
	replaceValues(root.Right, depth+1, sums)
}

func getSums(node *TreeNode, depth int, result []int) []int {
	if node == nil {
		return result
	}

	if depth == len(result) {
		result = append(result, 0)
	}

	result[depth] += node.Val

	result = getSums(node.Left, depth+1, result)
	result = getSums(node.Right, depth+1, result)
	return result
}
