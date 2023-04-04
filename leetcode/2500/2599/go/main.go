package main

func makePrefSumNonNegative(nums []int) (result int) {
	heap := make([]int, 1, 64)

	for i, sum := 0, 0; i < len(nums); i++ {
		sum += nums[i]

		if nums[i] >= 0 {
			continue
		}

		if sum < 0 {
			sum -= replace(heap, nums[i])
			result++
		} else {
			heap = insert(heap, nums[i])
		}
	}

	return result
}

func replace(heap []int, x int) int {
	if heap[0] < x {
		x, heap[0] = heap[0], x
		siftDown(heap, 0)
	}

	return x
}

func insert(heap []int, x int) []int {
	heap = append(heap, x)
	siftUp(heap, len(heap)-1)

	return heap
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
