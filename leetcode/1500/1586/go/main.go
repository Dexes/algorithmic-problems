package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	lastNode *TreeNode

	stack      []*TreeNode
	stackIndex int

	list       []int
	listIndex  int
	listLength int
}

func Constructor(node *TreeNode) BSTIterator {
	return BSTIterator{lastNode: node, stack: dirtyStack, stackIndex: -1, list: dirtyList, listIndex: -1, listLength: 0}
}

func (x *BSTIterator) HasNext() bool {
	return x.lastNode != nil || x.stackIndex >= 0 || x.listIndex+1 < x.listLength
}

func (x *BSTIterator) Next() int {
	if x.listIndex++; x.listIndex < x.listLength {
		return x.list[x.listIndex]
	}

	for node := x.lastNode; node != nil; node = node.Left {
		x.stackIndex++
		x.stack[x.stackIndex] = node
	}

	x.lastNode = x.stack[x.stackIndex].Right
	x.list[x.listIndex] = x.stack[x.stackIndex].Val
	x.listLength++
	x.stackIndex--

	return x.list[x.listIndex]
}

func (x *BSTIterator) HasPrev() bool {
	return x.listIndex > 0
}

func (x *BSTIterator) Prev() int {
	x.listIndex--
	return x.list[x.listIndex]
}

var (
	dirtyStack = make([]*TreeNode, 25)
	dirtyList  = make([]int, 35000)
)
