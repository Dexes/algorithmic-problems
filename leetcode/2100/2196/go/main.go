package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	Node   *TreeNode
	IsRoot bool
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*NodeInfo)
	for i := 0; i < len(descriptions); i++ {
		parent, child := getNode(nodes, descriptions[i][0]), getNode(nodes, descriptions[i][1])
		child.IsRoot = false

		if descriptions[i][2] == 1 {
			parent.Node.Left = child.Node
		} else {
			parent.Node.Right = child.Node
		}
	}

	for _, info := range nodes {
		if info.IsRoot {
			return info.Node
		}
	}

	return nil
}

func getNode(nodes map[int]*NodeInfo, value int) *NodeInfo {
	if node, ok := nodes[value]; ok {
		return node
	}

	node := &NodeInfo{Node: &TreeNode{Val: value}, IsRoot: true}
	nodes[value] = node
	return node
}
