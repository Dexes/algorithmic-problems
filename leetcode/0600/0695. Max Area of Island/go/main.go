package main

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	queue := makeQueue()
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			area := getArea(queue, grid, i, j)
			if area > result {
				result = area
			}
		}
	}

	return result
}

func getArea(queue *Queue, grid [][]int, x, y int) int {
	result := 0
	queue.Push(x, y)
	for !queue.IsEmpty() {
		x, y = queue.Pop()
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
			continue
		}

		result++
		grid[x][y] = 0
		queue.Push(x+1, y)
		queue.Push(x-1, y)
		queue.Push(x, y+1)
		queue.Push(x, y-1)
	}

	return result
}

type Queue struct {
	x         []int
	y         []int
	pushIndex int
	popIndex  int
}

func makeQueue() *Queue {
	return &Queue{x: make([]int, 1000), y: make([]int, 1000), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(x, y int) {
	q.x[q.pushIndex] = x
	q.y[q.pushIndex] = y
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
