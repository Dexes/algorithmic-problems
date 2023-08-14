package main

func findKthLargest(nums []int, k int) int {
	heap := heapify(nums[:k])
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			siftDown(heap, 0)
		}
	}

	return heap[0]
}

func heapify(data []int) []int {
	for i := len(data)/2 - 1; i >= 0; i-- {
		siftDown(data, i)
	}

	return data
}

func siftDown(data []int, i int) {
	for cIndex := getChildIndex(data, i); cIndex > 0 && data[i] > data[cIndex]; cIndex = getChildIndex(data, i) {
		data[i], data[cIndex] = data[cIndex], data[i]
		i = cIndex
	}
}

func getChildIndex(data []int, i int) int {
	switch lIndex, rIndex := i*2+1, i*2+2; {
	case lIndex >= len(data):
		return -1
	case rIndex >= len(data):
		return lIndex
	case data[rIndex] > data[lIndex]:
		return lIndex
	default:
		return rIndex
	}
}
