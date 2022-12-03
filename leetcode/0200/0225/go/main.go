package main

type MyStack struct {
	first  *Queue
	second *Queue
}

func Constructor() MyStack {
	return MyStack{first: makeQueue(50), second: makeQueue(50)}
}

func (x *MyStack) Push(val int) {
	x.first.Push(val)
}

func (x *MyStack) Pop() int {
	x.transferData()
	return x.second.Pop()
}

func (x *MyStack) Top() int {
	x.transferData()

	result := x.second.Pop()
	x.first.Push(result)

	return result
}

func (x *MyStack) transferData() {
	for x.first.Len() > 1 {
		x.second.Push(x.first.Pop())
	}

	x.second, x.first = x.first, x.second
}

func (x *MyStack) Empty() bool {
	return x.first.IsEmpty() && x.second.IsEmpty()
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
