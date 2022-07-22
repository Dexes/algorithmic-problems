package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	leftHead, rightHead := &ListNode{}, &ListNode{}
	left, right := leftHead, rightHead

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}

		head = head.Next
	}

	left.Next = rightHead.Next
	right.Next = nil

	return leftHead.Next
}
