package main

func connectSticks(sticks []int) (result int) {
	if len(sticks) == 1 {
		return 0
	}

	for heap := NewHeap(sticks); heap.Len() > 1; {
		cost := heap.Extract() + heap.Get()
		heap.Replace(cost)

		result += cost
	}

	return result
}

type Heap struct {
	data []int
}

func NewHeap(data []int) *Heap {
	result := &Heap{data: data}
	for i := len(data)/2 - 1; i >= 0; i-- {
		result.siftDown(i)
	}

	return result
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

func (h *Heap) Get() int {
	return h.data[0]
}

func (h *Heap) Len() int {
	return len(h.data)
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
