package main

func findClosestNumber(nums []int) int {
	result, distance := nums[0], abs(nums[0])
	for i := 1; i < len(nums); i++ {
		if abs(nums[i]) < distance || nums[i] == distance {
			result, distance = nums[i], abs(nums[i])
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
