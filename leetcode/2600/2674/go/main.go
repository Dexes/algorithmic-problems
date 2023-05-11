package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitCircularLinkedList(first *ListNode) []*ListNode {
	tail, middle, shiftMiddle := first, first, true
	for ; tail.Next != first; tail = tail.Next {
		if shiftMiddle = !shiftMiddle; shiftMiddle {
			middle = middle.Next
		}
	}

	second := middle.Next

	tail.Next = second
	middle.Next = first

	return []*ListNode{first, second}
}
