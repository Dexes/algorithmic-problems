package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	head, tail := toList(root)
	head.Left, tail.Right = tail, head

	return head
}

func toList(root *Node) (*Node, *Node) {
	var headLeft, tailLeft, headRight, tailRight *Node = root, nil, nil, root
	if root.Left != nil {
		headLeft, tailLeft = toList(root.Left)
		root.Left, tailLeft.Right = tailLeft, root
	}

	if root.Right != nil {
		headRight, tailRight = toList(root.Right)
		root.Right, headRight.Left = headRight, root
	}

	return headLeft, tailRight
}
