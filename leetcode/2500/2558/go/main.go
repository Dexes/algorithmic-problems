package main

import "math"

func pickGifts(gifts []int, k int) int64 {
	heapify(gifts)

	for i := 0; i < k; i++ {
		if gifts[0] == 1 {
			return int64(len(gifts))
		}

		gifts[0] = int(math.Sqrt(float64(gifts[0])))
		siftDown(gifts, 0)
	}

	return sum(gifts)
}

func heapify(data []int) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		siftDown(data, i)
	}
}

func siftDown(data []int, i int) {
	for cIndex := getChildIndex(data, i); cIndex > 0 && data[i] < data[cIndex]; cIndex = getChildIndex(data, i) {
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
	case data[rIndex] < data[lIndex]:
		return lIndex
	default:
		return rIndex
	}
}

func sum(nums []int) int64 {
	for i := 1; i < len(nums); i++ {
		nums[0] += nums[i]
	}

	return int64(nums[0])
}
