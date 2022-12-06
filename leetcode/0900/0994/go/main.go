package main

func orangesRotting(grid [][]int) int {
	queue, freshNumber, maxMinute := makeQueue(100), 0, 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				freshNumber++
				continue
			}

			if grid[i][j] == 2 {
				queue.Push(QueueNode{i, j, 0})
			}
		}
	}

	for !queue.IsEmpty() {
		node := queue.Pop()
		maxMinute = node.minute + 1

		if x := node.i + 1; x < len(grid) && grid[x][node.j] == 1 {
			freshNumber--
			grid[x][node.j] = 2
			queue.Push(QueueNode{x, node.j, maxMinute})
		}

		if x := node.i - 1; x >= 0 && grid[x][node.j] == 1 {
			freshNumber--
			grid[x][node.j] = 2
			queue.Push(QueueNode{x, node.j, maxMinute})
		}

		if x := node.j + 1; x < len(grid[0]) && grid[node.i][x] == 1 {
			freshNumber--
			grid[node.i][x] = 2
			queue.Push(QueueNode{node.i, x, maxMinute})
		}

		if x := node.j - 1; x >= 0 && grid[node.i][x] == 1 {
			freshNumber--
			grid[node.i][x] = 2
			queue.Push(QueueNode{node.i, x, maxMinute})
		}
	}

	if freshNumber > 0 {
		return -1
	}

	return maxMinute - 1
}

type QueueNode struct {
	i, j, minute int
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
