package main

type RecentCounter struct {
	queue *Queue
}

func Constructor() RecentCounter {
	return RecentCounter{queue: makeQueue(3002)}
}

func (x *RecentCounter) Ping(t int) int {
	for start := t - 3000; !x.queue.IsEmpty() && x.queue.Top() < start; x.queue.Pop() {
	}

	x.queue.Push(t)
	return x.queue.Len()
}

type Queue struct {
	data      []int
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{data: make([]int, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(val int) {
	q.data[q.pushIndex] = val
	q.pushIndex++

	if q.pushIndex >= len(q.data) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() int {
	val := q.data[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.data) {
		q.popIndex = 0
	}

	return val
}

func (q *Queue) Top() int {
	return q.data[q.popIndex]
}

func (q *Queue) Len() int {
	result := q.pushIndex - q.popIndex
	if result < 0 {
		result += len(q.data)
	}

	return result
}

func (q *Queue) IsEmpty() bool {
	return q.pushIndex == q.popIndex
}
