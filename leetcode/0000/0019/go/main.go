package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var next, prev *ListNode
	current := &ListNode{0, &ListNode{0, head}}

	for current.Next != nil {
		if n == 1 {
			next = head
		} else if n < 1 {
			next = next.Next
		}

		if n == -1 {
			prev = head
		} else if n < -1 {
			prev = prev.Next
		}

		current = current.Next
		n--
	}

	if prev == nil {
		return head.Next
	}

	prev.Next = next
	return head
}
