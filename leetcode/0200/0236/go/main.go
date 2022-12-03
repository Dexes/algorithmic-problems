package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	var lca *TreeNode
	if lca = lowestCommonAncestor(root.Left, p, q); lca == nil {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if lowestCommonAncestor(root.Right, p, q) == nil {
		return lca
	}

	return root
}
