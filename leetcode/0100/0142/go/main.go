package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := floyd(head)
	if node == nil {
		return nil
	}

	for head != node {
		head, node = head.Next, node.Next
	}

	return head
}

func floyd(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return slow
		}
	}

	return nil
}
