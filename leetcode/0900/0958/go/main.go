package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var buffer = make([]*TreeNode, 100)

func isCompleteTree(root *TreeNode) bool {
	nodes, index := append(buffer[:0], root), 0
	for ; ; index++ {
		if nodes[index].Left == nil {
			if nodes[index].Right != nil {
				return false
			}

			return areLeaves(nodes[index+1:])
		}

		nodes = append(nodes, nodes[index].Left)

		if nodes[index].Right == nil {
			return areLeaves(nodes[index+1:])
		}

		nodes = append(nodes, nodes[index].Right)
	}
}

func areLeaves(nodes []*TreeNode) bool {
	for _, n := range nodes {
		if n.Left != nil || n.Right != nil {
			return false
		}
	}

	return true
}
