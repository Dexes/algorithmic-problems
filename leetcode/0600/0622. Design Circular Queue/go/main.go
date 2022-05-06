package main

func main() {
	x := Constructor(3)
	q := &x
	println(q.EnQueue(1))
	println(q.EnQueue(2))
	println(q.EnQueue(3))
	println(q.EnQueue(4))
	println(q.Rear())
	println(q.IsFull())
	println(q.DeQueue())
	println(q.EnQueue(4))
	println(q.Rear())
}

type MyCircularQueue struct {
	data      []int
	headIndex int
	tailIndex int
	length    int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{data: make([]int, k), headIndex: 0, tailIndex: -1, length: 0}
}

func (x *MyCircularQueue) EnQueue(value int) bool {
	if x.IsFull() {
		return false
	}

	x.tailIndex = x.incrementIndex(x.tailIndex)
	x.data[x.tailIndex] = value
	x.length++
	return true
}

func (x *MyCircularQueue) DeQueue() bool {
	if x.IsEmpty() {
		return false
	}

	x.headIndex = x.incrementIndex(x.headIndex)
	x.length--
	return true
}

func (x *MyCircularQueue) Front() int {
	if x.IsEmpty() {
		return -1
	}

	return x.data[x.headIndex]
}

func (x *MyCircularQueue) Rear() int {
	if x.IsEmpty() {
		return -1
	}

	return x.data[x.tailIndex]
}

func (x *MyCircularQueue) IsEmpty() bool {
	return x.length == 0
}

func (x *MyCircularQueue) IsFull() bool {
	return x.length == len(x.data)
}

func (x *MyCircularQueue) incrementIndex(index int) int {
	if index++; index < len(x.data) {
		return index
	}

	return 0
}
