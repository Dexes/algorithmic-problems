package main

func countNegatives(grid [][]int) int {
	result, left := 0, len(grid[0])-1
	for i := 0; i < len(grid); i++ {
		left = findLeft(grid[i], 0, left)

		if left == -1 {
			left = len(grid[i]) - 1
		} else {
			result += len(grid[i]) - left
		}
	}

	return result
}

func findLeft(nums []int, left, right int) int {
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] < 0 && (middle == 0 || nums[middle-1] >= 0) {
			return middle
		}

		if nums[middle] < 0 {
			right = middle
		} else {
			left = middle + 1
		}
	}

	if nums[right] < 0 {
		return right
	}

	return -1
}
