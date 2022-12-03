package main

func isConsecutive(nums []int) bool {
	min, index := getMin(nums), 0
	for i := 0; i < len(nums); {
		if index = nums[i] - min; index == i {
			i++
			continue
		}

		if index >= len(nums) || nums[index] == nums[i] {
			return false
		}

		nums[i], nums[index] = nums[index], nums[i]
	}

	return true
}

func getMin(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < result {
			result = nums[i]
		}
	}

	return result
}
