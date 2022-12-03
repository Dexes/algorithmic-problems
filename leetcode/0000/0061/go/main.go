package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length, tail := 0, head
	for current := head; current != nil; current, tail = current.Next, current {
		length++
	}

	if length < 2 {
		return head
	}

	if k %= length; k == 0 {
		return head
	}

	var prev, result *ListNode = nil, head
	for i := length - k; i > 0; i-- {
		prev, result = result, result.Next
	}

	prev.Next, tail.Next = nil, head
	return result
}
