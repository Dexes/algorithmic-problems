package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	for list := head; list.Next != nil; list = list.Next.Next {
		list.Next = &ListNode{Val: GCD(list.Val, list.Next.Val), Next: list.Next}
	}

	return head
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
