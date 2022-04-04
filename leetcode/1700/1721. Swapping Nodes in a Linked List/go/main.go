package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	var (
		first, second           *ListNode
		firstIndex, secondIndex = k, length(head) - k + 1
	)

	if firstIndex == secondIndex {
		return head
	}

	for index, current := 1, head; first == nil || second == nil; index++ {
		switch {
		case firstIndex == index:
			first = current
		case secondIndex == index:
			second = current
		}

		current = current.Next
	}

	first.Val, second.Val = second.Val, first.Val
	return head
}

func length(head *ListNode) (result int) {
	for result = 0; head != nil; result++ {
		head = head.Next
	}

	return result
}
