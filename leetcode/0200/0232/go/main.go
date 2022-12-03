package main

type MyQueue struct {
	first  *Stack
	second *Stack
}

func Constructor() MyQueue {
	return MyQueue{first: makeStack(50), second: makeStack(50)}
}

func (x *MyQueue) Push(val int) {
	x.first.Push(val)
}

func (x *MyQueue) Pop() int {
	if x.second.IsEmpty() {
		x.transferData()
	}

	return x.second.Pop()
}

func (x *MyQueue) Peek() int {
	if x.second.IsEmpty() {
		x.transferData()
	}

	return x.second.Top()
}

func (x *MyQueue) Empty() bool {
	return x.first.IsEmpty() && x.second.IsEmpty()
}

func (x *MyQueue) transferData() {
	for !x.first.IsEmpty() {
		x.second.Push(x.first.Pop())
	}
}

type Stack struct {
	data   []int
	index  int
	length int
}

func makeStack(capacity int) *Stack {
	return &Stack{data: make([]int, capacity), index: 0, length: 0}
}

func (x *Stack) Push(val int) {
	x.data[x.index] = val
	x.index++
	x.length++

	if x.index == len(x.data) {
		x.index = 0
	}
}

func (x *Stack) Pop() int {
	x.index--
	x.length--
	if x.index < 0 {
		x.index = 0
	}

	return x.data[x.index]
}

func (x *Stack) Top() int {
	index := x.index - 1
	if index < 0 {
		index = 0
	}

	return x.data[index]
}

func (x *Stack) IsEmpty() bool {
	return x.length == 0
}
