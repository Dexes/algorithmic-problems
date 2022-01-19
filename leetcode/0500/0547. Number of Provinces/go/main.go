package main

type Queue struct {
	data      []int
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{data: make([]int, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(n int) {
	q.data[q.pushIndex] = n
	q.pushIndex++

	if q.pushIndex >= len(q.data) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() int {
	result := q.data[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.data) {
		q.popIndex = 0
	}

	return result
}

func (q *Queue) IsEmpty() bool {
	return q.pushIndex == q.popIndex
}

func findCircleNum(isConnected [][]int) int {
	result := 0
	queue, visited := makeQueue(1000), make([]bool, len(isConnected))

	for i := 0; i < len(isConnected); i++ {
		if visited[i] {
			continue
		}

		queue.Push(i)
		for !queue.IsEmpty() {
			current := queue.Pop()
			visited[current] = true
			for j := i + 1; j < len(isConnected); j++ {
				if !visited[j] && isConnected[j][current] == 1 {
					queue.Push(j)
				}
			}
		}

		result++
	}

	return result
}
