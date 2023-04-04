package main

func maxSubArray(nums []int) int {
	current, result := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		current = max(current+nums[i], nums[i])
		result = max(result, current)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
