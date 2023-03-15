package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var queue = makeQueue(50)

func isCompleteTree(root *TreeNode) bool {
	queue.Clear()
	queue.Push(root)

	for {
		n := queue.Pop()
		lEmpty, rEmpty := n.Left == nil, n.Right == nil

		if lEmpty {
			if !rEmpty {
				return false
			}

			return checkTail(queue)
		}

		queue.Push(n.Left)

		if rEmpty {
			return checkTail(queue)
		}

		queue.Push(n.Right)
	}
}

func checkTail(q *Queue) bool {
	for !q.IsEmpty() {
		if n := q.Pop(); n.Left != nil || n.Right != nil {
			return false
		}
	}

	return true
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

func (q *Queue) Clear() {
	q.popIndex, q.pushIndex = 0, 0
}
