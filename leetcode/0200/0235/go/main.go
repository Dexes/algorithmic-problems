package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || (p.Val < root.Val && root.Val < q.Val) {
		return root
	}

	if q.Val < root.Val {
		return dfs(root.Left, p, q)
	}

	return dfs(root.Right, p, q)
}
