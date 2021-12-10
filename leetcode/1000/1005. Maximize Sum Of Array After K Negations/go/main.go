package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	k = inverseNegatives(nums, k)

	sum, min := 0, 10000
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}

	if k&1 == 1 {
		sum -= min * 2
	}

	return sum
}

func inverseNegatives(nums []int, k int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if k == 0 || nums[i] >= 0 {
			break
		}

		nums[i] = -nums[i]
		k--
	}

	return k
}
