package main

func findMiddleIndex(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	last := nums[len(nums)-1]
	if nums[0] == last {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == last-nums[i] {
			return i
		}
	}

	return -1
}
