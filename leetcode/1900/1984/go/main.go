package main

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	result := 100000
	for i := k - 1; i < len(nums); i++ {
		if nums[i]-nums[i-k+1] < result {
			result = nums[i] - nums[i-k+1]
		}
	}

	return result
}
