package main

import "sort"

func relocateMarbles(nums []int, from, to []int) []int {
	positions := make(map[int]struct{})
	for _, num := range nums {
		positions[num] = struct{}{}
	}

	for i := range from {
		if from[i] == to[i] {
			continue
		}

		positions[to[i]] = struct{}{}
		delete(positions, from[i])
	}

	nums = nums[:0]
	for p := range positions {
		nums = append(nums, p)
	}

	sort.Ints(nums)
	return nums
}
