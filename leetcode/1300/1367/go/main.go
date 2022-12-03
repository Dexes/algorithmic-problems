package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	return check(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func check(head *ListNode, root *TreeNode) bool {
	switch {
	case head == nil:
		return true
	case root == nil || head.Val != root.Val:
		return false
	default:
		return check(head.Next, root.Left) || check(head.Next, root.Right)
	}
}
