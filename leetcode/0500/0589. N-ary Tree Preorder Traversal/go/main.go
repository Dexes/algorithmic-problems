package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	return dfs(root, make([]int, 0, 5000))
}

func dfs(root *Node, result []int) []int {
	result = append(result, root.Val)

	for i := 0; i < len(root.Children); i++ {
		result = dfs(root.Children[i], result)
	}

	return result
}
