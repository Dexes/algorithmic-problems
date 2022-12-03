package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	numbers := make(map[int]bool)
	for i := 0; i < k && i < len(nums); i++ {
		if numbers[nums[i]] {
			return true
		}

		numbers[nums[i]] = true
	}

	for i := k; i < len(nums); i++ {
		if numbers[nums[i]] {
			return true
		}

		delete(numbers, nums[i-k])
		numbers[nums[i]] = true
	}

	return false
}
