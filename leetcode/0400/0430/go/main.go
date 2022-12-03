package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root != nil {
		unpack(root)
	}

	return root
}

func unpack(head *Node) (*Node, *Node) {
	for tail := head; ; tail = tail.Next {
		if tail.Child != nil {
			childHead, childTail := unpack(tail.Child)
			tail.Child = nil

			tail = insert(tail, childHead, childTail)
		}

		if tail.Next == nil {
			return head, tail
		}
	}
}

func insert(root, head, tail *Node) *Node {
	if next := root.Next; next != nil {
		tail.Next, next.Prev = next, tail
	}

	root.Next, head.Prev = head, root
	return tail
}
