package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head.Next == nil || head.Next.Next == nil {
		return
	}

	current, middle := head.Next, head
	for current != nil && current.Next != nil {
		middle, current = middle.Next, current.Next.Next
	}

	middle.Next, middle = nil, reverse(middle.Next)
	for current = head; middle != nil; current = current.Next.Next {
		current.Next, middle.Next, middle = middle, current.Next, middle.Next
	}
}

func reverse(head *ListNode) *ListNode {
	prev, current := head, head.Next
	head.Next = nil

	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}

	return prev
}
