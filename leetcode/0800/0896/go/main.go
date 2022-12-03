package main

func isMonotonic(nums []int) bool {
	i := 1
	for ; i < len(nums) && nums[i-1] == nums[i]; i++ {
	}

	if i == len(nums) {
		return true
	}

	increasing := nums[i-1] < nums[i]
	for i++; i < len(nums); i++ {
		if increasing && nums[i-1] > nums[i] {
			return false
		}

		if !increasing && nums[i-1] < nums[i] {
			return false
		}
	}

	return true
}
