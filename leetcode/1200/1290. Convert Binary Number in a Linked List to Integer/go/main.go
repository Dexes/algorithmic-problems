package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result <<= 1
		if head.Val == 1 {
			result |= 1
		}

		head = head.Next
	}

	return result
}
