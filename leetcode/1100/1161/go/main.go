package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue := makeQueue(1500)
	queue.Push(root.Left)
	queue.Push(root.Right)
	result, resultLevel, currentLevel := root.Val, 1, 2

	for !queue.IsEmpty() {
		levelSum := 0
		for i := queue.Len(); i > 0; i-- {
			node := queue.Pop()
			queue.Push(node.Left)
			queue.Push(node.Right)
			levelSum += node.Val
		}

		if result < levelSum {
			result, resultLevel = levelSum, currentLevel
		}

		currentLevel++
	}

	return resultLevel
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
