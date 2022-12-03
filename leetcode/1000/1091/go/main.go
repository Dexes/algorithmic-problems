package main

var directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	lastIndex := len(grid) - 1
	if grid[0][0] == 1 || grid[lastIndex][lastIndex] == 1 {
		return -1
	}

	var distance, x, y, nx, ny int
	queue := makeQueue(200)
	queue.Push(0, 0)
	grid[0][0] = 1

	for !queue.IsEmpty() {
		x, y = queue.Pop()
		distance = grid[x][y] + 1

		for _, d := range directions {
			nx, ny = x+d[0], y+d[1]
			if nx < 0 || nx > lastIndex || ny < 0 || ny > lastIndex || grid[nx][ny] != 0 {
				continue
			}

			grid[nx][ny] = distance
			queue.Push(nx, ny)
		}
	}

	if grid[lastIndex][lastIndex] == 0 {
		return -1
	}

	return grid[lastIndex][lastIndex]
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
