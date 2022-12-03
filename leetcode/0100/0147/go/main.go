package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	preHead := &ListNode{}
	for ; head != nil; head = head.Next {
		insert(preHead, head.Val)
	}

	return preHead.Next
}

func insert(preHead *ListNode, x int) {
	prev, current := preHead, preHead.Next

	for ; current != nil; prev, current = current, current.Next {
		if current.Val > x {
			prev.Next = &ListNode{Val: x, Next: current}
			return
		}
	}

	prev.Next = &ListNode{Val: x}
}
