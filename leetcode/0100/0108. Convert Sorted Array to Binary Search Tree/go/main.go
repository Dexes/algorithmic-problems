package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	index := len(nums) / 2
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:index]),
		Right: sortedArrayToBST(nums[index+1:]),
	}
}
