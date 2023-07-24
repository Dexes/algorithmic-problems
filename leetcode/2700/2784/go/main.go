package main

func isGood(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	last := len(nums) - 1
	for i := 0; i < last; i++ {
		for expect := i + 1; nums[i] != expect; {
			if nums[i] >= len(nums) {
				return false
			}

			if target := nums[i] - 1; nums[target] != nums[i] {
				nums[target], nums[i] = nums[i], nums[target]
				continue
			}

			if nums[i] == last && nums[i] != nums[last] {
				nums[last], nums[i] = nums[i], nums[last]
				continue
			}

			return false
		}
	}

	return nums[last-1] == nums[last]
}
