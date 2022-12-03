package main

func dominantIndex(nums []int) int {
	index, first, second := 0, -1, -1

	for i := 0; i < len(nums); i++ {
		switch {
		case first < nums[i]:
			index, first, second = i, nums[i], first
		case second < nums[i]:
			second = nums[i]
		}
	}

	if second<<1 <= first {
		return index
	}

	return -1
}
