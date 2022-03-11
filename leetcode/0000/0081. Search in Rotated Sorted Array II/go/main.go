package main

func search(nums []int, target int) bool {
	for left, right := 0, len(nums)-1; left <= right; {
		switch middle := left + (right-left)/2; {
		case nums[middle] == target:
			return true
		case nums[left] < nums[middle] || nums[right] < nums[middle]:
			if nums[middle] < target || target < nums[left] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		case nums[middle] < nums[right] || nums[middle] < nums[left]:
			if nums[right] < target || target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		default:
			left, right = left+1, right-1
		}
	}

	return false
}
