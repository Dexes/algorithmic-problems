package main

func minimumAverageDifference(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	rightIndex := len(nums) - 1
	min, index := difference(nums, 0, rightIndex), 0
	for i := 1; i < len(nums); i++ {
		if x := difference(nums, i, rightIndex); x < min {
			min, index = x, i
		}
	}

	return index
}

func difference(prefix []int, index, rightIndex int) int {
	left := prefix[index] / (index + 1)
	if rightIndex == index {
		return left
	}

	right := (prefix[rightIndex] - prefix[index]) / (rightIndex - index)
	if left > right {
		return left - right
	}

	return right - left
}
