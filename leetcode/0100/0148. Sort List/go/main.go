package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := cut(head)
	return merge(sortList(head), sortList(middle))
}

func cut(head *ListNode) *ListNode {
	preMiddle, middle, isMiddle := head, head, false

	for head.Next != nil {
		if isMiddle = !isMiddle; isMiddle {
			preMiddle, middle = middle, middle.Next
		}

		head = head.Next
	}

	preMiddle.Next = nil
	return middle
}

func merge(left *ListNode, right *ListNode) *ListNode {
	var result, current *ListNode

	result, left, right = getNode(left, right)
	current = result

	for left != nil || right != nil {
		current.Next, left, right = getNode(left, right)
		current = current.Next
	}

	return result
}

func getNode(left *ListNode, right *ListNode) (*ListNode, *ListNode, *ListNode) {
	if left == nil {
		if right == nil {
			return nil, nil, nil
		}

		return right, left, right.Next
	}

	if right == nil {
		return left, left.Next, right
	}

	if right.Val < left.Val {
		return right, left, right.Next
	}

	return left, left.Next, right
}
