package main

import "sort"

func minPairSum(nums []int) (result int) {
	sort.Ints(nums)

	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		if x := nums[l] + nums[r]; x > result {
			result = x
		}
	}

	return result
}
