package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		firstHead, secondHead = &ListNode{}, &ListNode{}
		first, second         = firstHead, secondHead
	)

	for current := head; ; current = current.Next.Next {
		first.Next, first = current, current
		second.Next, second = current.Next, current.Next
		if current.Next == nil || current.Next.Next == nil {
			break
		}
	}

	first.Next = secondHead.Next
	return firstHead.Next
}
