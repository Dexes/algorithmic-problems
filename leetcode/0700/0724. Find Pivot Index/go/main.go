package main

func pivotIndex(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	rightIndex := len(nums) - 1
	if nums[rightIndex]-nums[0] == 0 {
		return 0
	}

	for i := 1; i < rightIndex; i++ {
		if nums[i-1] == nums[rightIndex]-nums[i] {
			return i
		}
	}

	if nums[rightIndex-1] == 0 {
		return rightIndex
	}

	return -1
}
