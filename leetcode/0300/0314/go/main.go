package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	left, right := treeRange(root, 0, 0, 0)
	result, queue := make([][]int, right-left+1), makeQueue(20)
	queue.Push(QueueNode{Tree: root, Column: 0})

	for !queue.IsEmpty() {
		node := queue.Pop()
		insert(result, left, node)

		if node.Tree.Left != nil {
			queue.Push(QueueNode{Tree: node.Tree.Left, Column: node.Column - 1})
		}

		if node.Tree.Right != nil {
			queue.Push(QueueNode{Tree: node.Tree.Right, Column: node.Column + 1})
		}
	}

	return result
}

func insert(result [][]int, shift int, node QueueNode) {
	if node.Column -= shift; result[node.Column] == nil {
		result[node.Column] = make([]int, 0, 4)
	}

	result[node.Column] = append(result[node.Column], node.Tree.Val)
}

func treeRange(root *TreeNode, current, left, right int) (int, int) {
	if root.Left != nil {
		next := current - 1
		l, r := treeRange(root.Left, next, min(next, left), right)
		left, right = min(left, l), max(right, r)
	}

	if root.Right != nil {
		next := current + 1
		l, r := treeRange(root.Right, next, left, max(next, right))
		left, right = min(left, l), max(right, r)
	}

	return left, right
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type QueueNode struct {
	Tree   *TreeNode
	Column int
}

type Queue struct {
	nodes     []QueueNode
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{nodes: make([]QueueNode, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(node QueueNode) {
	q.nodes[q.pushIndex] = node
	q.pushIndex++

	if q.pushIndex >= len(q.nodes) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() QueueNode {
	node := q.nodes[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.nodes) {
		q.popIndex = 0
	}

	return node
}

func (q *Queue) IsEmpty() bool {
	return q.pushIndex == q.popIndex
}
