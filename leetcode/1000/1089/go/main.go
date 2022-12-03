package main

var queue = makeQueue(2000)

func duplicateZeros(arr []int) {
	queue.Clear()

	for i := 0; i < len(arr); i++ {
		if !queue.IsEmpty() {
			queue.Push(arr[i])
			arr[i] = queue.Pop()
		}

		if arr[i] == 0 && i < len(arr)-1 {
			i++
			queue.Push(arr[i])
			arr[i] = 0
		}
	}
}

type Queue struct {
	numbers   []int
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{numbers: make([]int, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(num int) {
	q.numbers[q.pushIndex] = num
	q.pushIndex++

	if q.pushIndex >= len(q.numbers) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() int {
	num := q.numbers[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.numbers) {
		q.popIndex = 0
	}

	return num
}

func (q *Queue) Clear() {
	q.pushIndex, q.popIndex = 0, 0
}

func (q *Queue) IsEmpty() bool {
	return q.pushIndex == q.popIndex
}
