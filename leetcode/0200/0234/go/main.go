package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	tail := reverse(cutInHalf(head))

	for head != nil {
		if head.Val != tail.Val {
			return false
		}

		head, tail = head.Next, tail.Next
	}

	return true
}

func cutInHalf(head *ListNode) *ListNode {
	prev, middle, isMiddle := head, head, false

	for head.Next != nil {
		if isMiddle = !isMiddle; isMiddle {
			prev, middle = middle, middle.Next
		}

		head = head.Next
	}

	prev.Next = nil
	return middle
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev, current := head, head.Next
	head.Next = nil

	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}

	return prev
}
