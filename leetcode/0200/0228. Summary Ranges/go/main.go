package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := make([]string, 0, len(nums))
	for i, j := 0, 0; i < len(nums); i = j {
		for j = i + 1; j < len(nums) && nums[j] == nums[j-1]+1; j++ {
		}

		result = append(result, toString(nums[i], nums[j-1]))
	}

	return result
}

func toString(i, j int) string {
	if i == j {
		return strconv.Itoa(i)
	}

	return strconv.Itoa(i) + "->" + strconv.Itoa(j)
}
