package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := left + (right-left)/2

		if target <= nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return left
}
