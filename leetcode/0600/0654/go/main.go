package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	result := &TreeNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		insert(result, nums[i])
	}

	return result
}

func insert(node *TreeNode, val int) {
	for {
		if node.Val < val {
			*node = TreeNode{Val: val, Left: &TreeNode{Val: node.Val, Left: node.Left, Right: node.Right}}
			return
		}

		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
			return
		}

		node = node.Right
	}
}
