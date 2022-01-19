package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	result, queue := make([][]int, 0, 500), makeQueue(3500)
	queue.Push(root)

	for !queue.IsEmpty() {
		level := make([]int, queue.Len())
		for i := 0; i < len(level); i++ {
			node := queue.Pop()
			level[i] = node.Val

			for j := 0; j < len(node.Children); j++ {
				queue.Push(node.Children[j])
			}
		}

		result = append(result, level)
	}

	return result
}

type Queue struct {
	nodes     []*Node
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{nodes: make([]*Node, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(node *Node) {
	q.nodes[q.pushIndex] = node
	q.pushIndex++

	if q.pushIndex >= len(q.nodes) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() *Node {
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
