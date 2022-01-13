package main

func containsDuplicate(nums []int) bool {
	numbers := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if numbers[nums[i]] {
			return true
		}

		numbers[nums[i]] = true
	}

	return false
}
