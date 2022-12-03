package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack      []*TreeNode
	stackIndex int
}

func Constructor(node *TreeNode) BSTIterator {
	x := &BSTIterator{stack: make([]*TreeNode, 100), stackIndex: -1}
	for node != nil {
		x.stackIndex++
		x.stack[x.stackIndex] = node
		node = node.Left
	}

	return *x
}

func (x *BSTIterator) Next() int {
	result := x.stack[x.stackIndex]
	x.stackIndex--

	for node := result.Right; node != nil; node = node.Left {
		x.stackIndex++
		x.stack[x.stackIndex] = node
	}

	return result.Val
}

func (x *BSTIterator) HasNext() bool {
	return x.stackIndex >= 0
}
