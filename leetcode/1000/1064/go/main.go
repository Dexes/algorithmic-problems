package main

func fixedPoint(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		if middle := left + (right-left)/2; arr[middle] < middle {
			left = middle + 1
		} else {
			right = middle
		}
	}

	if arr[left] == left {
		return left
	}

	return -1
}
