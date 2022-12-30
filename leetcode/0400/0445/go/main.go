package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(head1, head2 *ListNode) *ListNode {
	if head1.Val == 0 {
		return head2
	}

	if head2.Val == 0 {
		return head1
	}

	tail1, tail2 := reverse(head1), reverse(head2)
	add(tail1, tail2)

	return reverse(tail1)
}

func add(tail1, tail2 *ListNode) {
	var (
		prev1 *ListNode
		carry bool
	)

	for tail1 != nil && tail2 != nil {
		if carry {
			tail1.Val, carry = tail1.Val+1, false
		}

		if tail1.Val += tail2.Val; tail1.Val >= 10 {
			tail1.Val, carry = tail1.Val-10, true
		}

		prev1, tail1, tail2 = tail1, tail1.Next, tail2.Next
	}

	if tail1 == nil && tail2 == nil {
		if carry {
			prev1.Next = &ListNode{Val: 1}
		}

		return
	}

	if tail1 == nil && tail2 != nil {
		prev1.Next, tail1 = tail2, tail2
	}

	if !carry {
		return
	}

	for {
		if tail1 == nil {
			prev1.Next = &ListNode{Val: 1}
			return
		}

		if tail1.Val += 1; tail1.Val < 10 {
			return
		}

		tail1.Val -= 10
		prev1, tail1 = tail1, tail1.Next
	}
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
