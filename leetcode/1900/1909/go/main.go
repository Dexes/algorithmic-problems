package main

func canBeIncreasing(nums []int) bool {
	first := true
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			continue
		}

		if !first {
			return false
		}

		first = false

		if i > 1 && nums[i-2] >= nums[i] {
			nums[i] = nums[i-1]
		}
	}

	return true
}
