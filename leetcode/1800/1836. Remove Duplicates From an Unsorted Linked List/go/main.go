package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	set := make([]int, 1e5+1)

	for current := head; current != nil; current = current.Next {
		set[current.Val]++
	}

	result := &ListNode{Next: head}
	for prev, current := result, head; current != nil; current = current.Next {
		if set[current.Val] > 1 {
			prev.Next = current.Next
		} else {
			prev = current
		}
	}

	return result.Next
}
