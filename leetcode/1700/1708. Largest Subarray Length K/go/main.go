package main

func largestSubarray(nums []int, k int) []int {
	max, maxIndex := nums[0], 0
	for i := len(nums) - k; i > 0; i-- {
		if nums[i] > max {
			max, maxIndex = nums[i], i
		}
	}

	return nums[maxIndex : maxIndex+k]
}
