package main

func findMissingRanges(nums []int, lower, upper int) [][]int {
	if len(nums) == 0 {
		return [][]int{{lower, upper}}
	}

	result := make([][]int, 0, len(nums)+1)
	if nums[0]-lower > 0 {
		result = append(result, []int{lower, nums[0] - 1})
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			result = append(result, []int{nums[i-1] + 1, nums[i] - 1})
		}
	}

	if i := len(nums) - 1; upper-nums[i] > 0 {
		result = append(result, []int{nums[i] + 1, upper})
	}

	return result
}
