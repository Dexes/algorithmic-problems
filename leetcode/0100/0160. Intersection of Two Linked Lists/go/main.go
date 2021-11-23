package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodes := make(map[*ListNode]bool)
	for headA != nil {
		nodes[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if nodes[headB] {
			return headB
		}

		headB = headB.Next
	}

	return nil
}
