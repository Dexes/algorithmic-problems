package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	result, queue := make([]float64, 0, 1000), makeQueue()
	queue.Push(root)

	for !queue.IsEmpty() {
		levelLength, sum := queue.Len(), 0
		for i := 0; i < levelLength; i++ {
			node := queue.Pop()
			queue.Push(node.Left)
			queue.Push(node.Right)

			sum += node.Val
		}

		result = append(result, float64(sum)/float64(levelLength))
	}

	return result
}

type Queue struct {
	nodes     []*TreeNode
	pushIndex int
	popIndex  int
}

func makeQueue() *Queue {
	return &Queue{nodes: make([]*TreeNode, 1000), pushIndex: 0, popIndex: 0}
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
