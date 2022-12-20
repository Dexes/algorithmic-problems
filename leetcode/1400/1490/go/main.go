package main

type Node struct {
	Val      int
	Children []*Node
}

func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	return clone(root)
}

func clone(root *Node) *Node {
	result := &Node{Val: root.Val, Children: make([]*Node, len(root.Children))}
	for i, child := range root.Children {
		result.Children[i] = clone(child)
	}

	return result
}
