package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev, current := dummy, head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}

			prev.Next, current = current.Next, current.Next
		} else {
			prev, current = current, current.Next
		}
	}

	return dummy.Next
}
