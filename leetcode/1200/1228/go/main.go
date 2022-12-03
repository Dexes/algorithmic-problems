package main

func missingNumber(nums []int) int {
	step := getProgressionStep(nums)
	if step == 0 {
		return nums[0]
	}

	for i := 1; ; i++ {
		if nums[i]-nums[i-1] != step {
			return nums[i] - step
		}
	}
}

func getProgressionStep(nums []int) int {
	a, b := nums[1]-nums[0], nums[2]-nums[1]
	if (a < 0 && a > b) || a < b {
		return a
	}

	return b
}
