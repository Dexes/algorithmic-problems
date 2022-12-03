package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	nodes := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := nodes[head]; ok {
			return true
		}

		nodes[head] = struct{}{}
		head = head.Next
	}

	return false
}
