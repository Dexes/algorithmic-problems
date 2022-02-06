package main

func removeDuplicates(nums []int) int {
	index, current, currentCount := 1, nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == current {
			currentCount++
		} else {
			current, currentCount = nums[i], 1
		}

		if currentCount <= 2 {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
