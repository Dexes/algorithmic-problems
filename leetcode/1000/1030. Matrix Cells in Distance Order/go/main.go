package main

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	result := make([][]int, 0, rows*cols)
	checker, queue := makeAccessChecker(rows, cols), makeQueue(1000)
	checker.HasAccess(rCenter, cCenter)
	queue.Push(rCenter, cCenter)

	for !queue.IsEmpty() {
		x, y := queue.Pop()
		result = append(result, []int{x, y})

		if checker.HasAccess(x+1, y) {
			queue.Push(x+1, y)
		}

		if checker.HasAccess(x-1, y) {
			queue.Push(x-1, y)
		}

		if checker.HasAccess(x, y+1) {
			queue.Push(x, y+1)
		}

		if checker.HasAccess(x, y-1) {
			queue.Push(x, y-1)
		}
	}

	return result
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

type AccessChecker struct {
	visited    map[ACKey]struct{}
	rows, cols int
}

type ACKey struct {
	x, y int
}

func makeAccessChecker(rows, cols int) *AccessChecker {
	return &AccessChecker{visited: make(map[ACKey]struct{}), rows: rows, cols: cols}
}

func (a *AccessChecker) HasAccess(x, y int) bool {
	if x < 0 || x >= a.rows || y < 0 || y >= a.cols {
		return false
	}

	key := ACKey{x, y}
	if _, ok := a.visited[key]; ok {
		return false
	}

	a.visited[key] = struct{}{}
	return true
}
