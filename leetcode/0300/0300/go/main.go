package main

import "sort"

func lengthOfLIS(nums []int) int {
	subsequenceLength := 0

	for _, n := range nums {
		i := sort.Search(subsequenceLength, func(i int) bool { return nums[i] >= n })
		if i < subsequenceLength {
			nums[i] = n
			continue
		}

		nums[subsequenceLength] = n
		subsequenceLength++
	}

	return subsequenceLength
}
