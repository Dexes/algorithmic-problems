package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	lastNode   *TreeNode
	stack      []*TreeNode
	stackIndex int
}

func Constructor(node *TreeNode) BSTIterator {
	return BSTIterator{lastNode: node, stack: make([]*TreeNode, 100), stackIndex: -1}
}

func (x *BSTIterator) Next() int {
	for node := x.lastNode; node != nil; node = node.Left {
		x.stackIndex++
		x.stack[x.stackIndex] = node
	}

	x.lastNode = x.stack[x.stackIndex].Right
	result := x.stack[x.stackIndex].Val
	x.stackIndex--

	return result
}

func (x *BSTIterator) HasNext() bool {
	return x.lastNode != nil || x.stackIndex >= 0
}
