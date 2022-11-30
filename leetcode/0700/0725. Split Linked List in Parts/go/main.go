package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var (
		length       = getLength(head)
		tail         *ListNode
		result       = make([]*ListNode, 0, k)
		chunkSize, x = length / k, length % k
	)

	for head != nil {
		if x > 0 {
			tail, x = getTail(head, chunkSize+1), x-1
		} else {
			tail = getTail(head, chunkSize)
		}

		result = append(result, head)
		head, tail.Next = tail.Next, nil
	}

	for ; len(result) < cap(result); result = append(result, nil) {
	}

	return result
}

func getTail(head *ListNode, length int) *ListNode {
	for i := 1; i < length; i, head = i+1, head.Next {
	}

	return head
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}

	return length
}
