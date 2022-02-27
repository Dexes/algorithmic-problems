package main

func islandPerimeter(grid [][]int) int {
	queue := makeQueue(1000)
	queue.Push(findStartPoint(grid))
	n, m := len(grid), len(grid[0])
	result := 0

	for !queue.IsEmpty() {
		x, y := queue.Pop()
		if grid[x][y] == -1 {
			continue
		}

		grid[x][y] = -1

		if x == 0 || grid[x-1][y] == 0 {
			result++
		} else if x > 0 && grid[x-1][y] == 1 {
			queue.Push(x-1, y)
		}

		if x == n-1 || grid[x+1][y] == 0 {
			result++
		} else if x < n-1 && grid[x+1][y] == 1 {
			queue.Push(x+1, y)
		}

		if y == 0 || grid[x][y-1] == 0 {
			result++
		} else if y > 0 && grid[x][y-1] == 1 {
			queue.Push(x, y-1)
		}

		if y == m-1 || grid[x][y+1] == 0 {
			result++
		} else if y < m-1 && grid[x][y+1] == 1 {
			queue.Push(x, y+1)
		}
	}

	return result
}

func findStartPoint(grid [][]int) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return i, j
			}
		}
	}

	return -1, -1
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
