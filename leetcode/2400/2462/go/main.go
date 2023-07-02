package main

import (
	"sort"
)

func totalCost(costs []int, k int, candidates int) int64 {
	if len(costs) == k {
		return int64(sum(costs))
	}

	if candidates*2 >= len(costs) {
		sort.Ints(costs)
		return int64(sum(costs[:k]))
	}

	lIndex, rIndex := candidates, len(costs)-candidates-1
	lHeap, rHeap := heapify(costs[:lIndex]), heapify(costs[rIndex+1:])
	result := 0

	for i := 0; i < k; i++ {
		if len(rHeap) == 0 || (len(lHeap) > 0 && lHeap[0] <= rHeap[0]) {
			if lIndex <= rIndex {
				result += replace(lHeap, costs[lIndex])
				lIndex++
			} else {
				result += lHeap[0]
				lHeap = remove(lHeap)
			}
		} else {
			if lIndex <= rIndex {
				result += replace(rHeap, costs[rIndex])
				rIndex--
			} else {
				result += rHeap[0]
				rHeap = remove(rHeap)
			}
		}
	}

	return int64(result)
}

func replace(heap []int, val int) (result int) {
	result, heap[0] = heap[0], val
	siftDown(heap, 0)
	return result
}

func remove(heap []int) []int {
	lastIndex := len(heap) - 1
	heap[0] = heap[lastIndex]
	heap = heap[:lastIndex]
	siftDown(heap, 0)
	return heap
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
