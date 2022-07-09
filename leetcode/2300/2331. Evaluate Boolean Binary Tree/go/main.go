package _go

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const nodeFalse = 0
const nodeTrue = 1
const nodeOr = 2
const nodeAnd = 3

func evaluateTree(root *TreeNode) bool {
	switch {
	case root.Val == nodeFalse:
		return false
	case root.Val == nodeTrue:
		return true
	case root.Val == nodeOr:
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	default:
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}
}
