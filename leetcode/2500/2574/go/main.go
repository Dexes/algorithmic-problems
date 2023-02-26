package main

func leftRigthDifference(nums []int) []int {
	left, right := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		right += nums[i]
	}

	for i, num := range nums {
		right -= num
		nums[i] = abs(left - right)
		left += num
	}

	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
