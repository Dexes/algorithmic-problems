package main

import "sort"

func countElements(nums []int) int {
	sort.Ints(nums)

	lastIndex := len(nums) - 1
	if nums[0] == nums[lastIndex] {
		return 0
	}

	left, right := 1, lastIndex-1
	for ; nums[0] == nums[left]; left++ {
	}

	for ; nums[lastIndex] == nums[right]; right-- {
	}

	return right - left + 1
}
