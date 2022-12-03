package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	prev, current, index := list1, list1, 0
	for ; index < a; prev, current, index = current, current.Next, index+1 {
	}

	prev.Next = list2

	for ; list2.Next != nil; list2 = list2.Next {
	}

	for ; index < b; current, index = current.Next, index+1 {
	}

	list2.Next = current.Next

	return list1
}
