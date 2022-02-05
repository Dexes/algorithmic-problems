package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	result, resultIndex := make([]int, 5000), 0
	current, currentCount, maxCount := 0, 0, 0

	walkInorder(root, func(node *TreeNode) {
		if node.Val == current {
			currentCount++
		} else {
			current, currentCount = node.Val, 1
		}

		if maxCount < currentCount {
			result[0], resultIndex, maxCount = current, 1, currentCount
			return
		}

		if currentCount == maxCount {
			result[resultIndex] = current
			resultIndex++
		}
	})

	return result[:resultIndex]
}

func walkInorder(root *TreeNode, callback func(node *TreeNode)) {
	if root == nil {
		return
	}

	walkInorder(root.Left, callback)
	callback(root)
	walkInorder(root.Right, callback)
}
