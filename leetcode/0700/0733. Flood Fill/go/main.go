package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	queue, color := makeQueue(1000), image[sr][sc]
	n, m := len(image), len(image[0])
	queue.Push(sr, sc)
	for !queue.IsEmpty() {
		sr, sc = queue.Pop()
		if sr < 0 || sc < 0 || sr >= n || sc >= m || image[sr][sc] != color {
			continue
		}

		image[sr][sc] = newColor
		queue.Push(sr+1, sc)
		queue.Push(sr-1, sc)
		queue.Push(sr, sc+1)
		queue.Push(sr, sc-1)
	}

	return image
}

type Queue struct {
	x         []int
	y         []int
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{x: make([]int, capacity), y: make([]int, capacity), pushIndex: 0, popIndex: 0}
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
