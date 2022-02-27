package main

func numIslands(grid [][]byte) int {
	result := 0
	queue := makeQueue(1000)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}

			floodIsland(queue, grid, i, j)
			result++
		}
	}

	return result
}

func floodIsland(queue *Queue, grid [][]byte, x, y int) {
	queue.Push(x, y)
	for !queue.IsEmpty() {
		x, y = queue.Pop()
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
			continue
		}

		grid[x][y] = 0
		queue.Push(x+1, y)
		queue.Push(x-1, y)
		queue.Push(x, y+1)
		queue.Push(x, y-1)
	}
}

type Queue struct {
	x, y      []int
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{x: make([]int, capacity), y: make([]int, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(x, y int) {
	q.x[q.pushIndex], q.y[q.pushIndex] = x, y
	q.pushIndex++

	if q.pushIndex >= len(q.x) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() (int, int) {
	x, y := q.x[q.popIndex], q.y[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.x) {
		q.popIndex = 0
	}

	return x, y
}

func (q *Queue) IsEmpty() bool {
	return q.pushIndex == q.popIndex
}
