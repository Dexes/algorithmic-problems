package main

var (
	directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue      = makeQueue(2000)
)

func shortestBridge(grid [][]int) int {
	queue.Clear()

	i, j := searchIsland(grid)
	pushIsland(queue, grid, i, j)

	for distance := 0; ; distance++ {
		if expand(queue, grid) {
			return distance
		}
	}
}

func expand(q *Queue, grid [][]int) bool {
	var i, j, ni, nj int

	for x := queue.Len(); x > 0; x-- {
		i, j = queue.Pop()

		for _, d := range directions {
			if ni, nj = i+d[0], j+d[1]; !valid(grid, ni, nj) || grid[ni][nj] < 0 {
				continue
			}

			if grid[ni][nj] == 1 {
				return true
			}

			grid[ni][nj] = -1
			q.Push(ni, nj)
		}
	}

	return false
}

func searchIsland(grid [][]int) (int, int) {
	for i := 0; ; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return i, j
			}
		}
	}
}

func pushIsland(q *Queue, grid [][]int, i, j int) {
	if !valid(grid, i, j) || grid[i][j] < 1 {
		return
	}

	grid[i][j] = -1
	q.Push(i, j)

	pushIsland(q, grid, i-1, j)
	pushIsland(q, grid, i+1, j)
	pushIsland(q, grid, i, j-1)
	pushIsland(q, grid, i, j+1)
}

func valid(grid [][]int, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[i])
}

type Queue struct {
	rows, cols []int
	pushIndex  int
	popIndex   int
}

func makeQueue(capacity int) *Queue {
	return &Queue{rows: make([]int, capacity), cols: make([]int, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Clear() {
	q.popIndex, q.pushIndex = 0, 0
}

func (q *Queue) Push(row, col int) {
	q.rows[q.pushIndex], q.cols[q.pushIndex] = row, col
	q.pushIndex++

	if q.pushIndex >= len(q.rows) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() (int, int) {
	row, col := q.rows[q.popIndex], q.cols[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.rows) {
		q.popIndex = 0
	}

	return row, col
}

func (q *Queue) Len() int {
	result := q.pushIndex - q.popIndex
	if result < 0 {
		result += len(q.rows)
	}

	return result
}
