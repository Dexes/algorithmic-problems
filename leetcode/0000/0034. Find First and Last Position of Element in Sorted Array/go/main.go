package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1

	if left = searchLeft(nums, left, right, target); left == -1 {
		return []int{-1, -1}
	}

	return []int{left, searchRight(nums, left, right, target)}
}

func searchLeft(nums []int, left, right, target int) int {
	for left < right {
		if middle := left + (right-left)/2; nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func searchRight(nums []int, left, right, target int) int {
	for left < right {
		if middle := left + (right-left)/2 + 1; nums[middle] > target {
			right = middle - 1
		} else {
			left = middle
		}
	}

	if nums[right] == target {
		return right
	}

	return -1
}
