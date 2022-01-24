package main

import "math"

func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > first {
			first, second, third = nums[i], first, second
			continue
		}

		if nums[i] > second && nums[i] != first {
			second, third = nums[i], second
			continue
		}

		if nums[i] > third && nums[i] < second {
			third = nums[i]
		}
	}

	if third == math.MinInt64 {
		return first
	}

	return third
}
