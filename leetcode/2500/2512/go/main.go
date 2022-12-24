package main

func topStudents(positiveFeedbacks []string, negativeFeedbacks []string, reports []string, studentIDs []int, k int) []int {
	positive, negative := toSet(positiveFeedbacks), toSet(negativeFeedbacks)
	heap := NewHeap(k)

	for i := 0; i < len(reports); i++ {
		points := getPoints(reports[i], positive, negative)
		if heap.Len() < k {
			heap.Insert(HeapNode{StudentID: studentIDs[i], Points: points})
			continue
		}

		if min := heap.Get(); (min.Points == points && min.StudentID > studentIDs[i]) || min.Points < points {
			heap.Replace(HeapNode{StudentID: studentIDs[i], Points: points})
		}
	}

	result := make([]int, heap.Len())
	for i := heap.Len() - 1; i >= 0; i-- {
		result[i] = heap.Extract().StudentID
	}

	return result
}

func getPoints(report string, positive, negative map[string]struct{}) (result int) {
	for start, end := 0, 0; start < len(report); start = end + 1 {
		for end = start + 1; end < len(report) && report[end] != ' '; end++ {
		}

		if _, ok := positive[report[start:end]]; ok {
			result += 3
			continue
		}

		if _, ok := negative[report[start:end]]; ok {
			result -= 1
		}
	}

	return result
}

func toSet(arr []string) map[string]struct{} {
	result := make(map[string]struct{})
	for i := 0; i < len(arr); i++ {
		result[arr[i]] = struct{}{}
	}

	return result
}

type HeapNode struct {
	StudentID int
	Points    int
}

type Heap struct {
	data []HeapNode
}

func NewHeap(capacity int) *Heap {
	return &Heap{data: make([]HeapNode, 0, capacity)}
}

func (h *Heap) Insert(n HeapNode) {
	h.data = append(h.data, n)
	h.siftUp(len(h.data) - 1)
}

func (h *Heap) Extract() HeapNode {
	result, lastIndex := h.data[0], len(h.data)-1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	h.siftDown(0)

	return result
}

func (h *Heap) Replace(n HeapNode) (result HeapNode) {
	result, h.data[0] = h.data[0], n
	h.siftDown(0)

	return result
}

func (h *Heap) Get() HeapNode {
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
	return (h.data[a].Points == h.data[b].Points && h.data[a].StudentID < h.data[b].StudentID) || h.data[a].Points > h.data[b].Points
}
