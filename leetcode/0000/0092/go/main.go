package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if left == right {
		return head
	}

	head = &ListNode{Next: head}
	preMiddle, middle := split(head, left)
	_, tail := split(middle, right-left+1)

	preMiddle.Next, middle.Next = reverse(middle), tail

	return head.Next
}

func split(head *ListNode, index int) (*ListNode, *ListNode) {
	var prev, node *ListNode = nil, head
	for i := 0; i < index; i++ {
		prev, node = node, node.Next
	}

	prev.Next = nil
	return prev, node
}

func reverse(head *ListNode) *ListNode {
	prev, current := head, head.Next
	head.Next = nil

	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}

	return prev
}
