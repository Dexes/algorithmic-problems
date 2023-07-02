package main

func findMissingRanges(nums []int, lower, upper int) [][]int {
	if len(nums) == 0 {
		return [][]int{{lower - 1, upper + 1}}
	}

	result := make([][]int, 0, len(nums)+1)
	if nums[0]-lower > 0 {
		result = append(result, []int{lower - 1, nums[0]})
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			result = append(result, []int{nums[i-1], nums[i]})
		}
	}

	if i := len(nums) - 1; upper-nums[i] > 0 {
		result = append(result, []int{nums[i], upper + 1})
	}

	return result
}
