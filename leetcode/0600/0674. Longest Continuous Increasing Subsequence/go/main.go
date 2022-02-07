package main

func findLengthOfLCIS(nums []int) int {
	result, current := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			current++
			continue
		}

		if current > result {
			result = current
		}

		current = 1
	}

	if current > result {
		return current
	}

	return result
}
