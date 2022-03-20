package main

import "container/heap"

type FreqStack struct {
	counter int
	numbers map[int]int
	heap    heap.Interface
}

func Constructor() FreqStack {
	return FreqStack{counter: 0, numbers: make(map[int]int), heap: makeHeap(1000)}
}

func (x *FreqStack) Push(val int) {
	weight := x.numbers[val]
	heap.Push(x.heap, HeapNode{Number: val, Weight: weight, Order: x.counter})

	x.numbers[val] = weight + 1
	x.counter++
}

func (x *FreqStack) Pop() int {
	node := heap.Pop(x.heap).(HeapNode)
	x.numbers[node.Number]--
	return node.Number
}

type HeapNode struct {
	Number int
	Weight int
	Order  int
}

type Heap []HeapNode

func makeHeap(capacity int) *Heap {
	result := make(Heap, 0, capacity)
	return &result
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i].Weight == h[j].Weight {
		return h[i].Order > h[j].Order
	}

	return h[i].Weight > h[j].Weight
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(HeapNode))
}

func (h *Heap) Pop() interface{} {
	length := len(*h) - 1
	x := (*h)[length]
	*h = (*h)[:length]
	return x
}
