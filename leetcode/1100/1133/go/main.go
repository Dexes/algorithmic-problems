package main

import "sort"

func largestUniqueNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	rightIndex := len(nums) - 1
	if nums[rightIndex] != nums[rightIndex-1] {
		return nums[rightIndex]
	}

	for i := rightIndex - 1; i > 0; i-- {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	if nums[0] != nums[1] {
		return nums[0]
	}

	return -1
}
