package main

func semiOrderedPermutation(nums []int) int {
	minIndex, maxIndex := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[minIndex] {
			minIndex = i
		}

		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	result := minIndex + len(nums) - 1 - maxIndex
	if minIndex > maxIndex {
		result--
	}

	return result
}
