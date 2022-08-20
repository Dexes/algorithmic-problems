package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root, _ := makeNode(preorder, 0, 1001)
	return root
}

func makeNode(preorder []int, min, max int) (*TreeNode, []int) {
	if len(preorder) == 0 || preorder[0] < min || max < preorder[0] {
		return nil, preorder
	}

	root := TreeNode{Val: preorder[0]}
	root.Left, preorder = makeNode(preorder[1:], min, root.Val)
	root.Right, preorder = makeNode(preorder, root.Val, max)

	return &root, preorder
}
