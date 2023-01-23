package main

import "sort"

func maxScore(nums1, nums2 []int, k int) int64 {
	nums := toNums(nums1, nums2)
	sort.Slice(nums, func(i, j int) bool { return nums[i].Val > nums[j].Val })

	sum, heap := 0, heapify(nums[:k])
	for i := 0; i < k; i++ {
		sum += nums[i].Val
	}

	result := sum * heap[0].Weight
	for i := k; i < len(nums); i++ {
		sum += nums[i].Val - heap[0].Val

		heap[0] = nums[i]
		siftDown(heap, 0)

		if x := sum * heap[0].Weight; x > result {
			result = x
		}
	}

	return int64(result)
}

func heapify(data []Num) []Num {
	for i := len(data)/2 - 1; i >= 0; i-- {
		siftDown(data, i)
	}

	return data
}

func siftDown(data []Num, i int) {
	for cIndex := getChildIndex(data, i); cIndex > 0 && data[i].Weight > data[cIndex].Weight; cIndex = getChildIndex(data, i) {
		data[i], data[cIndex] = data[cIndex], data[i]
		i = cIndex
	}
}

func getChildIndex(data []Num, i int) int {
	switch lIndex, rIndex := i*2+1, i*2+2; {
	case lIndex >= len(data):
		return -1
	case rIndex >= len(data):
		return lIndex
	case data[rIndex].Weight > data[lIndex].Weight:
		return lIndex
	default:
		return rIndex
	}
}

func toNums(nums1, nums2 []int) []Num {
	result := make([]Num, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i].Val, result[i].Weight = nums1[i], nums2[i]
	}

	return result
}

type Num struct {
	Val, Weight int
}
