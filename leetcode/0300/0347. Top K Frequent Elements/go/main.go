package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	pq := makeHeap(k)

	for number, weight := range countNumbers(nums) {
		if pq.Len() < k {
			heap.Push(pq, HeapNode{Number: number, Weight: weight})
			continue
		}

		if pq.Top().Weight < weight {
			heap.Pop(pq)
			heap.Push(pq, HeapNode{Number: number, Weight: weight})
		}
	}

	return pq.Data()
}

func countNumbers(nums []int) map[int]int {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	return counter
}

type HeapNode struct {
	Number int
	Weight int
}

type Heap []HeapNode

func makeHeap(capacity int) *Heap {
	result := make(Heap, 0, capacity)
	return &result
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Top() HeapNode {
	return h[0]
}

func (h Heap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
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

func (h *Heap) Data() []int {
	pq, result := *h, make([]int, h.Len())
	for i := 0; i < len(result); i++ {
		result[i] = pq[i].Number
	}

	return result
}
