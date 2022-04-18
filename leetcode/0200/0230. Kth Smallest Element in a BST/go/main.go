package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	result, _ := inorder(root, 0, k)
	return result.Val
}

func inorder(root *TreeNode, size, k int) (*TreeNode, int) {
	if root == nil {
		return nil, size
	}

	result, size := inorder(root.Left, size, k)
	if result != nil {
		return result, 0
	}

	if size+1 == k {
		return root, 0
	}

	return inorder(root.Right, size+1, k)
}
