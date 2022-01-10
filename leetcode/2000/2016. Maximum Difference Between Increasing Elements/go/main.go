package main

func maximumDifference(nums []int) int {
	min, result := nums[0], -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			continue
		}

		if nums[i] > min && nums[i]-min > result {
			result = nums[i] - min
		}
	}

	return result
}
