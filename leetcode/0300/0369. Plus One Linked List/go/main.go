package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func plusOne(head *ListNode) *ListNode {
	var node *ListNode
	for current := head; current != nil; current = current.Next {
		if current.Val < 9 {
			node = current
		}
	}

	if node == nil {
		head = &ListNode{Val: 1, Next: head}
		node = head
	} else {
		node.Val++
	}

	for node = node.Next; node != nil; node = node.Next {
		node.Val = 0
	}

	return head
}
