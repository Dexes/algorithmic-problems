package main

type SmallestInfiniteSet struct {
	min  int
	heap []int
	set  [1001]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{min: 0, heap: make([]int, 0, 1000)}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if len(s.heap) == 0 {
		s.min++
		return s.min
	}

	last, result := len(s.heap)-1, s.heap[0]
	s.heap[0], s.heap = s.heap[last], s.heap[:last]
	siftDown(s.heap, 0)

	s.set[result] = false

	return result
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num > s.min || s.set[num] {
		return
	}

	s.set[num] = true

	s.heap = append(s.heap, num)
	siftUp(s.heap, len(s.heap)-1)
}

func siftDown(heap []int, i int) {
	for cIndex := getChildIndex(heap, i); cIndex > 0 && heap[i] > heap[cIndex]; cIndex = getChildIndex(heap, i) {
		heap[i], heap[cIndex] = heap[cIndex], heap[i]
		i = cIndex
	}
}

func siftUp(heap []int, i int) {
	for pIndex := (i - 1) / 2; heap[pIndex] > heap[i]; pIndex = (i - 1) / 2 {
		heap[i], heap[pIndex] = heap[pIndex], heap[i]
		i = pIndex
	}
}

func getChildIndex(heap []int, i int) int {
	switch lIndex, rIndex := i*2+1, i*2+2; {
	case lIndex >= len(heap):
		return -1
	case rIndex >= len(heap):
		return lIndex
	case heap[rIndex] > heap[lIndex]:
		return lIndex
	default:
		return rIndex
	}
}
