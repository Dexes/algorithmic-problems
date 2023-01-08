package main

import "sort"

func maximumCount(nums []int) int {
	pos := len(nums) - sort.Search(len(nums), func(i int) bool { return nums[i] > 0 })
	neg := sort.Search(len(nums)-pos, func(i int) bool { return nums[i] >= 0 })

	if neg > pos {
		return neg
	}

	return pos
}
