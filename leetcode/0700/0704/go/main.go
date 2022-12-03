package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)/2

		if nums[middle] == target {
			return middle
		}

		if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
