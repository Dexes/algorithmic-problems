package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortLinkedList(head *ListNode) *ListNode {
	for prev, current := head, head.Next; current != nil; {
		if current.Val >= 0 {
			prev, current = current, current.Next
			continue
		}

		next := current.Next
		prev.Next, current.Next, head = current.Next, head, current
		current = next
	}

	return head
}
