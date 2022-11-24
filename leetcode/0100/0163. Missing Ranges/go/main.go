package main

import (
	"strconv"
)

func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		return []string{makeRange(lower-1, upper+1)}
	}

	result := make([]string, 0, len(nums)+1)
	if nums[0]-lower > 0 {
		result = append(result, makeRange(lower-1, nums[0]))
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			result = append(result, makeRange(nums[i-1], nums[i]))
		}
	}

	if i := len(nums) - 1; upper-nums[i] > 0 {
		result = append(result, makeRange(nums[i], upper+1))
	}

	return result
}

func makeRange(a, b int) string {
	if b-a == 2 {
		return strconv.Itoa(a + 1)
	}

	return strconv.Itoa(a+1) + "->" + strconv.Itoa(b-1)
}
