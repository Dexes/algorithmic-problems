package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	preMiddle, middle := middleNode(head)
	preMiddle.Next = nil

	return &TreeNode{
		Val:   middle.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(middle.Next),
	}
}

func middleNode(head *ListNode) (*ListNode, *ListNode) {
	preMiddle, middle, isMiddle := head, head, false
	for head.Next != nil {
		if isMiddle = !isMiddle; isMiddle {
			preMiddle, middle = middle, middle.Next
		}

		head = head.Next
	}

	return preMiddle, middle
}
