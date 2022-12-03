package main

type NumberContainers struct {
	indices map[int]*Heap
	numbers map[int]*Number
}

func Constructor() NumberContainers {
	return NumberContainers{
		indices: make(map[int]*Heap),
		numbers: make(map[int]*Number),
	}
}

func (x NumberContainers) Change(index int, value int) {
	var (
		heap   *Heap
		number *Number
		exists bool
	)

	if number, exists = x.numbers[index]; exists {
		if number.value == value {
			return
		}

		x.indices[number.value].Remove(number)
		number.value = value
	} else {
		number = &Number{value: value, index: index}
		x.numbers[index] = number
	}

	if heap, exists = x.indices[value]; !exists {
		heap = NewHeap(4)
		x.indices[value] = heap
	}

	heap.Insert(number)
}

func (x NumberContainers) Find(number int) int {
	if heap, exists := x.indices[number]; exists {
		return heap.Get()
	}

	return -1
}

type Number struct {
	index     int
	value     int
	heapIndex int
}

type Heap struct {
	data []*Number
}

func NewHeap(capacity int) *Heap {
	return &Heap{data: make([]*Number, 0, capacity)}
}

func (h *Heap) Insert(x *Number) {
	x.heapIndex = len(h.data)

	h.data = append(h.data, x)
	h.siftUp(x.heapIndex)
}

func (h *Heap) Get() int {
	if len(h.data) == 0 {
		return -1
	}

	return h.data[0].index
}

func (h *Heap) Remove(x *Number) {
	removedIndex, lastIndex := x.heapIndex, len(h.data)-1
	h.swap(removedIndex, lastIndex)
	h.data = h.data[:lastIndex]

	if removedIndex != lastIndex {
		if pIndex := (removedIndex - 1) / 2; !h.less(pIndex, removedIndex) {
			h.siftDown(removedIndex)
		} else {
			h.siftUp(removedIndex)
		}
	}
}

func (h Heap) siftUp(i int) {
	for pIndex := (i - 1) / 2; h.less(pIndex, i); pIndex = (i - 1) / 2 {
		h.swap(pIndex, i)
		i = pIndex
	}
}

func (h Heap) siftDown(i int) {
	for cIndex := h.getChildIndex(i); cIndex > 0 && h.less(i, cIndex); cIndex = h.getChildIndex(i) {
		h.swap(i, cIndex)
		i = cIndex
	}
}

func (h Heap) getChildIndex(i int) int {
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

func (h Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].heapIndex, h.data[j].heapIndex = h.data[j].heapIndex, h.data[i].heapIndex
}

func (h Heap) less(i, j int) bool {
	return h.data[i].index > h.data[j].index
}
