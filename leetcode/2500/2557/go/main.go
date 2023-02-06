package main

import "sort"

func maxCount(banned []int, n int, maxSum64 int64) (result int) {
	sort.Ints(banned)

	tailIndex := sort.Search(len(banned), func(i int) bool { return banned[i] > n })
	banned = banned[:tailIndex]
	maxSum := int(maxSum64)

	bIndex, sum := 0, 0
	for i := 1; i <= n; i++ {
		if bIndex < len(banned) && i == banned[bIndex] {
			i, bIndex = skip(banned, bIndex)
			continue
		}

		if sum += i; sum > maxSum {
			return result
		}

		result++
	}

	return result
}

func skip(nums []int, index int) (int, int) {
	var prev int

	for {
		prev, index = nums[index], index+1
		if index == len(nums) || prev+1 <= nums[index] {
			return prev, index
		}
	}
}
