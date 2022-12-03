package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	result := &Node{Val: head.Val}
	nodes, currentSrc, currentDst := map[*Node]*Node{head: result}, head.Next, result

	for currentSrc != nil {
		node := &Node{Val: currentSrc.Val}
		nodes[currentSrc] = node

		currentDst.Next, currentDst = node, node
		currentSrc = currentSrc.Next
	}

	currentSrc, currentDst = head, result
	for currentSrc != nil {
		currentDst.Random = nodes[currentSrc.Random]
		currentSrc, currentDst = currentSrc.Next, currentDst.Next
	}

	return result
}
