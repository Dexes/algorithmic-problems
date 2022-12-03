package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result <<= 1
		result |= head.Val

		head = head.Next
	}

	return result
}
