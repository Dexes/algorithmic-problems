package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) (result int) {
	first, second := reverse(middleNode(head)), head
	for ; first != nil; first, second = first.Next, second.Next {
		if x := first.Val + second.Val; x > result {
			result = x
		}
	}

	return result
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

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev, current := head, head.Next
	head.Next = nil

	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}

	return prev
}
