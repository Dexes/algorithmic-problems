package main

import "sort"

func missingElement(nums []int, k int) int {
	i := sort.Search(len(nums), func(i int) bool { return nums[i]-nums[0]-i >= k }) - 1
	return nums[0] + k + i
}
