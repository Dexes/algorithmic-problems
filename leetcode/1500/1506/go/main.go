package main

type Node struct {
	Val      int
	Children []*Node
}

func findRoot(tree []*Node) *Node {
	mask := 0
	for _, parent := range tree {
		mask ^= parent.Val
		for _, child := range parent.Children {
			mask ^= child.Val
		}
	}

	for _, n := range tree {
		if n.Val == mask {
			return n
		}
	}

	return nil
}
