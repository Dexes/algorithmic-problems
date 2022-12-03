package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	currentNode := result
	buffer := 0

	for {
		first, second := getNodeValue(l1), getNodeValue(l2)
		currentNode.Val, buffer = plus(first, second, buffer)

		l1, l2 = getNextNode(l1), getNextNode(l2)
		if l1 == nil && l2 == nil {
			if buffer > 0 {
				currentNode.Next = &ListNode{Val: buffer}
			}

			return result
		}

		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}
}

func plus(first, second, buffer int) (int, int) {
	result := first + second + buffer
	if result < 10 {
		return result, 0
	}

	return result % 10, 1
}

func getNextNode(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	return node.Next
}

func getNodeValue(node *ListNode) int {
	if node == nil {
		return 0
	}

	return node.Val
}
