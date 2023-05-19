package main

type RopeTreeNode struct {
	len   int
	val   string
	left  *RopeTreeNode
	right *RopeTreeNode
}

func getKthCharacter(root *RopeTreeNode, k int) byte {
	var l int

	for root.len > 0 {
		if l = length(root.left); l >= k {
			root = root.left
		} else {
			k -= l
			root = root.right
		}
	}

	return root.val[k-1]
}

func length(node *RopeTreeNode) int {
	if node == nil {
		return 0
	}

	if node.len == 0 {
		return len(node.val)
	}

	return node.len
}
