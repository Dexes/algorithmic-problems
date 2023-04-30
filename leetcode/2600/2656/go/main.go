package main

func maximizeSum(nums []int, k int) int {
	return (k * max(nums)) + (k * (k - 1) / 2)
}

func max(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}
