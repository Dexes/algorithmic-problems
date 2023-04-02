package main

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	if k == len(reward1) {
		return sum(reward1)
	}

	if k == 0 {
		return sum(reward2)
	}

	for i := 0; i < len(reward1); i++ {
		reward1[i] -= reward2[i]
	}

	heap := heapify(reward1[:k])
	for i := k; i < len(reward1); i++ {
		if heap[0] >= reward1[i] {
			continue
		}

		heap[0] = reward1[i]
		siftDown(heap, 0)
	}

	return sum(reward2) + sum(heap)
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

func sum(nums []int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}

	return result
}
