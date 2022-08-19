package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	_, _, result := countAverage(root)
	return result
}

func countAverage(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}

	leftSum, leftCount, leftResult := countAverage(root.Left)
	rightSum, rightCount, rightResult := countAverage(root.Right)

	sum := leftSum + rightSum + root.Val
	count := leftCount + rightCount + 1
	result := leftResult + rightResult

	if sum/count == root.Val {
		result++
	}

	return sum, count, result
}
