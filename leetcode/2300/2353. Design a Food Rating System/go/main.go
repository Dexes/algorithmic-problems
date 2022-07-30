package main

type FoodRatings struct {
	cuisinesHeaps map[string]*Heap
	foodsHeaps    map[string]*Heap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	result := FoodRatings{}
	result.cuisinesHeaps = make(map[string]*Heap)
	result.foodsHeaps = make(map[string]*Heap, len(foods))

	var heap *Heap
	for i := 0; i < len(foods); i++ {
		if heap = result.cuisinesHeaps[cuisines[i]]; heap == nil {
			heap = NewHeap(len(foods))
			result.cuisinesHeaps[cuisines[i]] = heap
		}

		result.foodsHeaps[foods[i]] = heap
		heap.Insert(foods[i], ratings[i])
	}

	return result
}

func (x *FoodRatings) ChangeRating(food string, newRating int) {
	x.foodsHeaps[food].ChangeRating(food, newRating)
}

func (x *FoodRatings) HighestRated(cuisine string) string {
	return x.cuisinesHeaps[cuisine].Top()
}

type HeapNode struct {
	Food   string
	Rating int
}

type Heap struct {
	data    []HeapNode
	indices map[string]int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		data:    make([]HeapNode, 0, capacity),
		indices: make(map[string]int, capacity),
	}
}

func (h *Heap) Insert(food string, rating int) {
	h.data = append(h.data, HeapNode{Food: food, Rating: rating})
	h.indices[food] = len(h.data) - 1

	h.siftUp(len(h.data) - 1)
}

func (h *Heap) Top() string {
	return h.data[0].Food
}

func (h Heap) ChangeRating(food string, rating int) {
	index := h.indices[food]
	rating, h.data[index].Rating = h.data[index].Rating, rating

	switch {
	case h.data[index].Rating > rating: // rate up
		h.siftUp(index)
	case h.data[index].Rating < rating: // rate down
		h.siftDown(index)
	}
}

func (h Heap) Get() HeapNode {
	return h.data[0]
}

func (h Heap) siftUp(i int) {
	for pIndex := (i - 1) / 2; h.less(pIndex, i); pIndex = (i - 1) / 2 {
		h.swap(i, pIndex)
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
	iFood, jFood := h.data[i].Food, h.data[j].Food

	h.indices[iFood], h.indices[jFood] = h.indices[jFood], h.indices[iFood]
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h Heap) less(a, b int) bool {
	if h.data[a].Rating == h.data[b].Rating {
		return h.data[a].Food > h.data[b].Food
	}

	return h.data[a].Rating < h.data[b].Rating
}
