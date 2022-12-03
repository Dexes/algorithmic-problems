package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	result, _ := findTiltAndSum(root)
	return result
}

func findTiltAndSum(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	tiltLeft, sumLeft := findTiltAndSum(root.Left)
	tiltRight, sumRight := findTiltAndSum(root.Right)

	return tiltLeft + tiltRight + abs(sumLeft-sumRight), sumLeft + sumRight + root.Val
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
