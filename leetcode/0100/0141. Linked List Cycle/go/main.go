package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	nodes := make(map[*ListNode]bool)
	for head != nil {
		if nodes[head] {
			return true
		}

		nodes[head] = true
		head = head.Next
	}

	return false
}
