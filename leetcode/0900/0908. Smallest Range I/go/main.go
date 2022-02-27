package main

func smallestRangeI(nums []int, k int) int {
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] > max:
			max = nums[i]
		case nums[i] < min:
			min = nums[i]
		}
	}

	min, max = min+k, max-k
	if min >= max {
		return 0
	}

	return max - min
}
