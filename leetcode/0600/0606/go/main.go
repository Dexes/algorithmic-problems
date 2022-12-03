package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	bytes := tree2bytes(root, make([]byte, 0, 50000))
	return string(bytes[1 : len(bytes)-1])
}

func tree2bytes(root *TreeNode, bytes []byte) []byte {
	if root == nil {
		return bytes
	}

	bytes = append(bytes, '(')
	bytes = append(bytes, strconv.Itoa(root.Val)...)

	if root.Left == nil && root.Right != nil {
		bytes = append(bytes, '(', ')')
	}

	bytes = tree2bytes(root.Left, bytes)
	bytes = tree2bytes(root.Right, bytes)

	bytes = append(bytes, ')')

	return bytes
}
