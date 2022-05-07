package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	read, write := head.Next, head
	for read != nil {
		write = write.Next
		write.Val, read = sum(read)
	}

	write.Next = nil
	return head.Next
}

func sum(head *ListNode) (int, *ListNode) {
	result := 0
	for ; head.Val != 0; head = head.Next {
		result += head.Val
	}

	return result, head.Next
}
