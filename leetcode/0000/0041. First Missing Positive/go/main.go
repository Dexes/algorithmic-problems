package main

func firstMissingPositive(nums []int) int {
	for rIndex := 0; rIndex < len(nums); rIndex++ {
		for {
			wIndex := nums[rIndex] - 1
			if nums[rIndex] <= 0 || nums[rIndex] > len(nums) || nums[wIndex] == wIndex+1 {
				break
			}

			nums[rIndex], nums[wIndex] = nums[wIndex], nums[rIndex]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
