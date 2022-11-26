package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	m, n = m-1, n+1

	for current := head; current != nil; current = current.Next {
		if current = skip(current, m); current == nil {
			break
		}

		current.Next = skip(current, n)
	}

	return head
}

func skip(head *ListNode, x int) *ListNode {
	for i := 0; i < x && head != nil; i, head = i+1, head.Next {
	}

	return head
}
