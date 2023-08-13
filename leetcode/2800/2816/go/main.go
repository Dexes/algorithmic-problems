package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	head = reverse(head)

	var (
		prev, current *ListNode = nil, head
		buffer                  = 0
	)

	for ; current != nil; prev, current = current, current.Next {
		if current.Val += current.Val + buffer; current.Val >= 10 {
			buffer = current.Val / 10
			current.Val = current.Val % 10
		} else {
			buffer = 0
		}
	}

	if buffer > 0 {
		prev.Next = &ListNode{Val: buffer}
	}

	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	prev, current := head, head.Next
	head.Next = nil
	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}

	return prev
}
