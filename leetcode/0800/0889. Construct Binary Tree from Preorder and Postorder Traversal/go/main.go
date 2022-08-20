package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) > 1 {
		i := search(postorder, preorder[1])
		root.Left = constructFromPrePost(preorder[1:i+2], postorder[:i+1])
		root.Right = constructFromPrePost(preorder[i+2:], postorder[i+1:])
	}

	return root
}

func search(data []int, target int) (result int) {
	for ; data[result] != target; result++ {
	}

	return result
}
