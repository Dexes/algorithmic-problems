package main

import "container/heap"

func topKFrequent(words []string, k int) []string {
	counter := make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		counter[words[i]]++
	}

	pq := makeHeap(k + 1)
	for word, frequency := range counter {
		heap.Push(pq, HeapNode{Word: word, Frequency: frequency})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	for i := k - 1; i >= 0; i-- {
		words[i] = heap.Pop(pq).(HeapNode).Word
	}

	return words[:k]
}

type HeapNode struct {
	Word      string
	Frequency int
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
	if h[i].Frequency == h[j].Frequency {
		return h[i].Word > h[j].Word
	}

	return h[i].Frequency < h[j].Frequency
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(HeapNode))
}

func (h *Heap) Pop() interface{} {
	right := len(*h) - 1
	x := (*h)[right]
	*h = (*h)[:right]
	return x
}
