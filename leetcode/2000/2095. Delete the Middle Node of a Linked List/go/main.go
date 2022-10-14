package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	preMiddle, middle := middleNode(head)
	preMiddle.Next = middle.Next
	return head
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
