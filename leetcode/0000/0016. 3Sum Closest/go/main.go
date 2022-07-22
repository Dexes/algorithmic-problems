package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	result := nums[0] + nums[1] + nums[2]
	diff := abs(target - result)

	for i := len(nums) - 1; i >= 3; i-- {
		for left, right := 0, i-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			currentDiff := abs(target - sum)

			if currentDiff == 0 {
				return sum
			}

			if currentDiff < diff {
				result, diff = sum, currentDiff
			}

			if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
