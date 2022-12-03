package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	if nums[0] >= len(nums) {
		return len(nums)
	}

	for j := len(nums) - 1; j > 1; j-- {
		if x := len(nums) - j; nums[j-1] < x && x <= nums[j] {
			return x
		}
	}

	return -1
}
