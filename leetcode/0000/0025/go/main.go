package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	result := &ListNode{}
	resultTail := result

	for tail, ok := cutNodes(head, k); ok; tail, ok = cutNodes(head, k) {
		resultTail.Next, resultTail = reverseList(head)
		head = tail
	}

	resultTail.Next = head
	return result.Next
}

func cutNodes(head *ListNode, k int) (*ListNode, bool) {
	if head == nil {
		return nil, false
	}

	prev, current := head, head.Next
	for i := 1; i < k; i++ {
		if current == nil {
			return nil, false
		}

		prev, current = current, current.Next
	}

	prev.Next = nil
	return current, true
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	tail, prev, current := head, head, head.Next
	head.Next = nil

	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}

	return prev, tail
}
