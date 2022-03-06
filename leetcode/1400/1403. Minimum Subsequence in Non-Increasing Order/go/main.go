package main

import "sort"

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})

	leftSum, rightSum, i := nums[0], 0, 1
	for ; i < len(nums); i++ {
		rightSum += nums[i]
	}

	for i = 1; leftSum <= rightSum; i++ {
		leftSum += nums[i]
		rightSum -= nums[i]
	}

	return nums[:i]
}
