package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(first *ListNode) *ListNode {
	if first == nil || first.Next == nil {
		return first
	}

	second := first.Next
	first.Next, second.Next = second.Next, first
	first.Next = swapPairs(first.Next)

	return second
}
