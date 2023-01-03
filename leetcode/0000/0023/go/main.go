package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	median := len(lists) / 2
	return mergeTwoLists(mergeKLists(lists[:median]), mergeKLists(lists[median:]))
}

func mergeTwoLists(left *ListNode, right *ListNode) *ListNode {
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
