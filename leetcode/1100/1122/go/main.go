package main

import (
	"sort"
)

func relativeSortArray(first []int, second []int) []int {
	indices := getIndices(second)
	sort.Slice(first, func(i, j int) bool {
		if indices[first[i]] > 0 && indices[first[j]] > 0 {
			return indices[first[i]] < indices[first[j]]
		}

		if indices[first[i]] > 0 {
			return true
		}

		if indices[first[j]] > 0 {
			return false
		}

		return first[i] < first[j]
	})

	return first
}

func getIndices(nums []int) []int {
	result := make([]int, 1001)
	for i := 0; i < len(nums); i++ {
		result[nums[i]] = i + 1
	}

	return result
}
