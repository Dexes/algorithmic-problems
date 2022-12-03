package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	queue := makeQueue(100)
	queue.Push(root)

	for !queue.IsEmpty() {
		xParent, yParent := inspectLevel(queue, x, y)
		if xParent != nil && yParent != nil && xParent != yParent {
			return true
		}

		if xParent != yParent {
			return false
		}
	}

	return false
}

func inspectLevel(queue *Queue, x, y int) (*TreeNode, *TreeNode) {
	var xParent, yParent *TreeNode

	length := queue.Len()
	for i := 0; i < length; i++ {
		node := queue.Pop()
		if childExists(node, x) {
			xParent = node
		}

		if childExists(node, y) {
			yParent = node
		}

		queue.Push(node.Left)
		queue.Push(node.Right)
	}

	return xParent, yParent
}

func childExists(node *TreeNode, x int) bool {
	if node == nil {
		return false
	}

	return node.Left != nil && node.Left.Val == x || node.Right != nil && node.Right.Val == x
}

type Queue struct {
	nodes     []*TreeNode
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{nodes: make([]*TreeNode, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(node *TreeNode) {
	if node == nil {
		return
	}

	q.nodes[q.pushIndex] = node
	q.pushIndex++

	if q.pushIndex >= len(q.nodes) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() *TreeNode {
	node := q.nodes[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.nodes) {
		q.popIndex = 0
	}

	return node
}

func (q *Queue) Len() int {
	result := q.pushIndex - q.popIndex
	if result < 0 {
		result += len(q.nodes)
	}

	return result
}

func (q *Queue) IsEmpty() bool {
	return q.pushIndex == q.popIndex
}
