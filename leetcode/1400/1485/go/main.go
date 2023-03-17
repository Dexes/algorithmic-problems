package main

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Random *Node
}

type NodeCopy struct {
	Val    int
	Left   *NodeCopy
	Right  *NodeCopy
	Random *NodeCopy
}

func copyRandomBinaryTree(root *Node) *NodeCopy {
	copies := make(map[*Node]*NodeCopy)
	copyNodes(root, copies)

	return copyTree(root, copies)
}

func copyTree(root *Node, copies map[*Node]*NodeCopy) *NodeCopy {
	if root == nil {
		return nil
	}

	result := copies[root]
	result.Left = copyTree(root.Left, copies)
	result.Right = copyTree(root.Right, copies)
	result.Random = copies[root.Random]

	return result
}

func copyNodes(root *Node, copies map[*Node]*NodeCopy) {
	if root == nil {
		return
	}

	copies[root] = &NodeCopy{Val: root.Val}
	copyNodes(root.Left, copies)
	copyNodes(root.Right, copies)
}
