package main

import "sort"

func singleNonDuplicate(nums []int) int {
	index := sort.Search(len(nums)/2, func(i int) bool {
		i *= 2
		return nums[i] != nums[i+1]
	}) * 2

	return nums[index]
}
