package main

func maxArrayValue(nums []int) int64 {
	result := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > result {
			result = nums[i]
			continue
		}

		result += nums[i]
	}

	return int64(result)
}
