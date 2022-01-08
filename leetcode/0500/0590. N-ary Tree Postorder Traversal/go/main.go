package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	return dfs(root, make([]int, 0, 5000))
}

func dfs(root *Node, result []int) []int {
	for i := 0; i < len(root.Children); i++ {
		result = dfs(root.Children[i], result)
	}

	return append(result, root.Val)
}
