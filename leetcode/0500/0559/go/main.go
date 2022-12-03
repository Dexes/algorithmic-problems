package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	result := 0
	for i := 0; i < len(root.Children); i++ {
		depth := maxDepth(root.Children[i])
		if depth > result {
			result = depth
		}
	}

	return result + 1
}
