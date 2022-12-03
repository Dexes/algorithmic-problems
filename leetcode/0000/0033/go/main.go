package main

func search(nums []int, target int) int {
	index := startIndex(nums)
	if target < nums[0] {
		return binarySearch(nums, index, len(nums)-1, target)
	}

	return binarySearch(nums, 0, index-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		switch middle := left + (right-left)/2; {
		case nums[middle] < target:
			left = middle + 1
		case nums[middle] > target:
			right = middle - 1
		default:
			return middle
		}
	}

	return -1
}

func startIndex(nums []int) int {
	for left, right := 1, len(nums)-1; left <= right; {
		switch middle := left + (right-left)/2; {
		case nums[0] < nums[middle]:
			left = middle + 1
		case nums[middle-1] < nums[middle]:
			right = middle - 1
		default:
			return middle
		}
	}

	return len(nums)
}
