package main

import (
	"container/heap"
)

func lastStoneWeight(stones []int) int {
	pq := toHeap(stones)
	for pq.Len() > 1 {
		first, second := heap.Pop(pq).(int), heap.Pop(pq).(int)
		if first > second {
			heap.Push(pq, first-second)
		}
	}

	if pq.Len() > 0 {
		return pq.Top()
	}

	return 0
}

func toHeap(data []int) *Heap {
	pq := Heap(data)
	heap.Init(&pq)
	return &pq
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Top() int {
	return h[0]
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	length := len(*h) - 1
	x := (*h)[length]
	*h = (*h)[:length]
	return x
}
