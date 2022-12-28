package main

func minStoneSum(piles []int, k int) (result int) {
	heapify(piles)

	for i := 0; i < k && piles[0] > 1; i++ {
		piles[0] = (piles[0] + 1) / 2
		siftDown(piles, 0)
	}

	for _, x := range piles {
		result += x
	}

	return result
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
