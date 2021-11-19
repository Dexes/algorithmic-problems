package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middle, isMiddle := head, false

	for head.Next != nil {
		if isMiddle = !isMiddle; isMiddle {
			middle = middle.Next
		}

		head = head.Next
	}

	return middle
}
