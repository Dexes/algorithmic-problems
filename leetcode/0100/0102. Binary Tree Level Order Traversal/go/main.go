package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result, queue := make([][]int, 0, 800), makeQueue(200)
	queue.Push(root)

	for !queue.IsEmpty() {
		level := make([]int, queue.Len())
		for i := 0; i < len(level); i++ {
			node := queue.Pop()
			level[i] = node.Val

			queue.Push(node.Left)
			queue.Push(node.Right)
		}

		result = append(result, level)
	}

	return result
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
