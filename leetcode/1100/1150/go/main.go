package main

func isMajorityElement(nums []int, target int) bool {
	var left, right int
	if left = searchLeft(nums, 0, len(nums)-1, target); left == -1 {
		return false
	}

	right = searchRight(nums, left, len(nums)-1, target)

	return right-left+1 > (len(nums) >> 1)
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
