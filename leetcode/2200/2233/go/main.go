package main

const modulo = 1e9 + 7

func maximumProduct(nums []int, k int) int {
	heapify(nums)
	for ; k > 0; k-- {
		nums[0]++
		siftDown(nums, 0)
	}

	if nums[0] == 0 {
		return 0
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result * nums[i] % modulo
	}

	return result
}

func heapify(data []int) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		siftDown(data, i)
	}
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
