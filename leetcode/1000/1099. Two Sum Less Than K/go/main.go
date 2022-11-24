package main

import "sort"

func twoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)
	result := -1

	for left, right := 0, len(nums)-1; left < right; {
		if sum := nums[left] + nums[right]; sum < k {
			result = max(result, sum)
			left++
		} else {
			right--
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
