package main

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)

	set := make(map[int]struct{})
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		set[nums[left]+nums[right]] = struct{}{}
	}

	return len(set)
}
