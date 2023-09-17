package main

func minimumRightShifts(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	index, lastIndex := 0, len(nums)-1
	for ; index < lastIndex && nums[index] < nums[index+1]; index++ {
	}

	if index == lastIndex {
		return 0
	}

	if nums[0] < nums[lastIndex] {
		return -1
	}

	index++
	result := len(nums) - index

	for ; index < lastIndex; index++ {
		if nums[index] > nums[index+1] {
			return -1
		}
	}

	return result
}
