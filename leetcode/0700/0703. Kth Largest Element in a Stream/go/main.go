package main

import "container/heap"

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Top() int {
	return h[0]
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
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

type KthLargest struct {
	Heap  *Heap
	Limit int
}

func Constructor(k int, nums []int) KthLargest {
	data := Heap(nums)
	heap.Init(&data)
	for data.Len() > k {
		heap.Pop(&data)
	}

	return KthLargest{Heap: &data, Limit: k}
}

func (x *KthLargest) Add(val int) int {
	if x.Heap.Len() < x.Limit {
		heap.Push(x.Heap, val)
	} else if val > x.Heap.Top() {
		heap.Pop(x.Heap)
		heap.Push(x.Heap, val)
	}

	return x.Heap.Top()
}
