package main

import "sort"

func highFive(items [][]int) [][]int {
	var (
		scores = make(map[int]*Heap)
		heap   *Heap
		ok     bool
		id     int
	)

	for i := 0; i < len(items); i++ {
		if heap, ok = scores[items[i][0]]; !ok {
			heap = NewHeap(5)
			scores[items[i][0]] = heap
		}

		if heap.Len() < 5 {
			heap.Insert(items[i][1])
			continue
		}

		if heap.Min() < items[i][1] {
			heap.Replace(items[i][1])
		}
	}

	result := make([][]int, 0, len(scores))
	for id, heap = range scores {
		result = append(result, []int{id, heap.Average()})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

type Heap struct {
	data []int
}

func NewHeap(capacity int) *Heap {
	return &Heap{data: make([]int, 0, capacity)}
}

func (h *Heap) Insert(n int) {
	h.data = append(h.data, n)
	h.siftUp(len(h.data) - 1)
}

func (h *Heap) Extract() int {
	result, lastIndex := h.data[0], len(h.data)-1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	h.siftDown(0)

	return result
}

func (h *Heap) Replace(n int) (result int) {
	result, h.data[0] = h.data[0], n
	h.siftDown(0)

	return result
}

func (h *Heap) Min() int {
	return h.data[0]
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Average() int {
	sum := 0
	for i := 0; i < len(h.data); i++ {
		sum += h.data[i]
	}

	return sum / len(h.data)
}

func (h *Heap) siftUp(i int) {
	for pIndex := (i - 1) / 2; h.less(pIndex, i); pIndex = (i - 1) / 2 {
		h.data[pIndex], h.data[i] = h.data[i], h.data[pIndex]
		i = pIndex
	}
}

func (h *Heap) siftDown(i int) {
	for cIndex := h.getChildIndex(i); cIndex > 0 && h.less(i, cIndex); cIndex = h.getChildIndex(i) {
		h.data[i], h.data[cIndex] = h.data[cIndex], h.data[i]
		i = cIndex
	}
}

func (h *Heap) getChildIndex(i int) int {
	lIndex, rIndex := i*2+1, i*2+2
	if lIndex >= len(h.data) {
		return -1
	}

	if rIndex >= len(h.data) {
		return lIndex
	}

	if h.less(rIndex, lIndex) {
		return lIndex
	}

	return rIndex
}

func (h *Heap) less(a, b int) bool {
	return h.data[a] > h.data[b]
}
