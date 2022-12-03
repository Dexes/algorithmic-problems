package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	head = reverse(head)

	for current := head; current != nil; current = current.Next {
		for ; current.Next != nil && current.Next.Val < current.Val; current.Next = current.Next.Next {
		}
	}

	return reverse(head)
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
