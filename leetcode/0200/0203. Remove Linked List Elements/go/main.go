package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Next.Val == val {
			current.Next = current.Next.Next
		}

		current = current.Next
	}

	if head != nil && head.Val == val {
		return head.Next
	}

	return head
}
