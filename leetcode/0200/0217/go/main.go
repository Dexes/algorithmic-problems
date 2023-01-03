package main

func containsDuplicate(nums []int) bool {
	numbers := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := numbers[nums[i]]; ok {
			return true
		}

		numbers[nums[i]] = struct{}{}
	}

	return false
}
